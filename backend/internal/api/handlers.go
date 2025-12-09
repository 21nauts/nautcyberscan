package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/nautcyberscanner/nautscanner/internal/osv"
	"github.com/nautcyberscanner/nautscanner/pkg/config"
	"github.com/nautcyberscanner/nautscanner/pkg/models"
)

// HealthCheck returns the API health status
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
		"service": "nautscanner-api",
	})
}

// ScanHandler handles scan-related requests
type ScanHandler struct {
	db  *sql.DB
	cfg *config.Config
}

func NewScanHandler(db *sql.DB, cfg *config.Config) *ScanHandler {
	return &ScanHandler{db: db, cfg: cfg}
}

// SubmitScan receives scan data from GitHub Actions
func (h *ScanHandler) SubmitScan(w http.ResponseWriter, r *http.Request) {
	var scanReq models.ScanRequest
	if err := json.NewDecoder(r.Body).Decode(&scanReq); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Create OSV client
	osvClient := osv.New()

	// Query OSV for vulnerabilities
	vulnerabilities, err := osvClient.BatchQuery(scanReq.Dependencies, scanReq.Ecosystem)
	if err != nil {
		// Log error but don't fail the scan
		vulnerabilities = []models.VulnerabilityMatch{}
	}

	// Save to database
	if err := h.saveScanResults(scanReq, vulnerabilities); err != nil {
		// Log error but continue - don't fail the scan
		fmt.Printf("Warning: failed to save scan results: %v\n", err)
	}

	// Create response
	response := models.ScanResponse{
		ScanID:               "scan-" + scanReq.RepoID,
		VulnerabilitiesFound: len(vulnerabilities),
		Vulnerabilities:      vulnerabilities,
		Timestamp:            time.Now(),
	}

	respondJSON(w, http.StatusOK, response)
}

func (h *ScanHandler) saveScanResults(scanReq models.ScanRequest, vulnerabilities []models.VulnerabilityMatch) error {
	// Start transaction
	tx, err := h.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback()

	// 1. Get or create default user (for now, use a default user)
	var userID string
	err = tx.QueryRow(`
		INSERT INTO users (email, name, plan)
		VALUES ('default@nautscanner.io', 'Default User', 'free')
		ON CONFLICT (email) DO UPDATE SET email = EXCLUDED.email
		RETURNING id
	`).Scan(&userID)
	if err != nil {
		return fmt.Errorf("failed to create/get user: %w", err)
	}

	// 2. Get or create repository
	var repoID string
	err = tx.QueryRow(`
		INSERT INTO repositories (user_id, name, full_name, last_scanned)
		VALUES ($1, $2, $3, NOW())
		ON CONFLICT (user_id, full_name)
		DO UPDATE SET last_scanned = NOW()
		RETURNING id
	`, userID, scanReq.RepoID, scanReq.RepoID).Scan(&repoID)
	if err != nil {
		return fmt.Errorf("failed to create/get repository: %w", err)
	}

	// 3. Save dependencies and vulnerabilities
	for _, dep := range scanReq.Dependencies {
		// Insert dependency
		var depID string
		err = tx.QueryRow(`
			INSERT INTO dependencies (repository_id, name, version, ecosystem, type)
			VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT (repository_id, name, ecosystem)
			DO UPDATE SET version = EXCLUDED.version
			RETURNING id
		`, repoID, dep.Name, dep.Version, scanReq.Ecosystem, dep.Type).Scan(&depID)
		if err != nil {
			return fmt.Errorf("failed to save dependency %s: %w", dep.Name, err)
		}

		// Find matching vulnerabilities for this dependency
		for _, vuln := range vulnerabilities {
			if vuln.DependencyName != dep.Name {
				continue
			}

			// Insert vulnerability
			var vulnID string
			err = tx.QueryRow(`
				INSERT INTO vulnerabilities (
					cve_id, title, description, severity, cvss_score,
					affected_package, ecosystem, patched_versions,
					published_at, exploited_at
				)
				VALUES ($1, $2, '', $3, $4, $5, $6, $7, NOW(), NULL)
				ON CONFLICT (cve_id)
				DO UPDATE SET
					title = EXCLUDED.title,
					severity = EXCLUDED.severity,
					cvss_score = EXCLUDED.cvss_score,
					updated_at = NOW()
				RETURNING id
			`, vuln.CVEID, vuln.Title, vuln.Severity, vuln.CVSSScore,
				vuln.DependencyName, scanReq.Ecosystem, vuln.PatchedVersions).Scan(&vulnID)
			if err != nil {
				return fmt.Errorf("failed to save vulnerability %s: %w", vuln.CVEID, err)
			}

			// Insert scan result
			_, err = tx.Exec(`
				INSERT INTO scan_results (repository_id, dependency_id, vulnerability_id, status, detected_at)
				VALUES ($1, $2, $3, 'vulnerable', NOW())
				ON CONFLICT (repository_id, dependency_id, vulnerability_id)
				DO UPDATE SET detected_at = NOW()
			`, repoID, depID, vulnID)
			if err != nil {
				return fmt.Errorf("failed to save scan result: %w", err)
			}
		}
	}

	return tx.Commit()
}

// RepositoryHandler handles repository-related requests
type RepositoryHandler struct {
	db *sql.DB
}

func NewRepositoryHandler(db *sql.DB) *RepositoryHandler {
	return &RepositoryHandler{db: db}
}

func (h *RepositoryHandler) ListRepositories(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT
			id, user_id, name, provider, full_name,
			private, last_scanned, created_at, updated_at
		FROM repositories
		ORDER BY last_scanned DESC NULLS LAST
		LIMIT 100
	`

	rows, err := h.db.Query(query)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to fetch repositories")
		return
	}
	defer rows.Close()

	var repositories []models.Repository
	for rows.Next() {
		var repo models.Repository
		err := rows.Scan(
			&repo.ID, &repo.UserID, &repo.Name, &repo.Provider,
			&repo.FullName, &repo.Private, &repo.LastScanned,
			&repo.CreatedAt, &repo.UpdatedAt,
		)
		if err != nil {
			respondError(w, http.StatusInternalServerError, "Failed to scan repositories")
			return
		}
		repositories = append(repositories, repo)
	}

	if repositories == nil {
		repositories = []models.Repository{}
	}

	respondJSON(w, http.StatusOK, repositories)
}

func (h *RepositoryHandler) GetRepository(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get repository
	respondJSON(w, http.StatusOK, map[string]string{"message": "Repository details"})
}

func (h *RepositoryHandler) GetRepositoryVulnerabilities(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement vulnerability listing for repository
	respondJSON(w, http.StatusOK, []models.VulnerabilityMatch{})
}

// VulnerabilityHandler handles vulnerability-related requests
type VulnerabilityHandler struct {
	db *sql.DB
}

func NewVulnerabilityHandler(db *sql.DB) *VulnerabilityHandler {
	return &VulnerabilityHandler{db: db}
}

func (h *VulnerabilityHandler) ListVulnerabilities(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT DISTINCT
			v.id, v.cve_id, v.title, COALESCE(v.description, '') as description,
			v.severity, COALESCE(v.cvss_score, 0) as cvss_score, v.affected_package,
			v.ecosystem, COALESCE(v.version_range, '') as version_range,
			COALESCE(v.patched_versions, '') as patched_versions,
			v.published_at, v.exploited_at, v.created_at, v.updated_at
		FROM vulnerabilities v
		INNER JOIN scan_results sr ON sr.vulnerability_id = v.id
		WHERE sr.status = 'vulnerable'
		ORDER BY v.created_at DESC
		LIMIT 100
	`

	rows, err := h.db.Query(query)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to fetch vulnerabilities")
		return
	}
	defer rows.Close()

	var vulnerabilities []models.Vulnerability
	for rows.Next() {
		var v models.Vulnerability
		err := rows.Scan(
			&v.ID, &v.CVEID, &v.Title, &v.Description,
			&v.Severity, &v.CVSSScore, &v.AffectedPackage,
			&v.Ecosystem, &v.VersionRange, &v.PatchedVersions,
			&v.PublishedAt, &v.ExploitedAt, &v.CreatedAt, &v.UpdatedAt,
		)
		if err != nil {
			fmt.Printf("Error scanning vulnerability row: %v\n", err)
			respondError(w, http.StatusInternalServerError, "Failed to scan vulnerabilities")
			return
		}
		vulnerabilities = append(vulnerabilities, v)
	}

	if vulnerabilities == nil {
		vulnerabilities = []models.Vulnerability{}
	}

	respondJSON(w, http.StatusOK, vulnerabilities)
}

func (h *VulnerabilityHandler) GetVulnerability(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get vulnerability
	respondJSON(w, http.StatusOK, map[string]string{"message": "Vulnerability details"})
}

// UserHandler handles user-related requests
type UserHandler struct {
	db *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get current user
	respondJSON(w, http.StatusOK, map[string]string{"message": "Current user"})
}

func (h *UserHandler) UpdateNotificationSettings(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement notification settings update
	respondJSON(w, http.StatusOK, map[string]string{"message": "Notification settings updated"})
}

// Helper functions
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}
