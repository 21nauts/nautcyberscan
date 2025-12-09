package api

import (
	"database/sql"
	"encoding/json"
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

	// Create response
	response := models.ScanResponse{
		ScanID:               "scan-" + scanReq.RepoID,
		VulnerabilitiesFound: len(vulnerabilities),
		Vulnerabilities:      vulnerabilities,
		Timestamp:            time.Now(),
	}

	// TODO: Save scan results to database

	respondJSON(w, http.StatusOK, response)
}

// RepositoryHandler handles repository-related requests
type RepositoryHandler struct {
	db *sql.DB
}

func NewRepositoryHandler(db *sql.DB) *RepositoryHandler {
	return &RepositoryHandler{db: db}
}

func (h *RepositoryHandler) ListRepositories(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement repository listing
	respondJSON(w, http.StatusOK, []models.Repository{})
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
	// TODO: Implement vulnerability listing
	respondJSON(w, http.StatusOK, []models.Vulnerability{})
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
