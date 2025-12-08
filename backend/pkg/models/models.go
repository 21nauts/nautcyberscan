package models

import "time"

// Repository represents a scanned repository
type Repository struct {
	ID          string    `json:"id" db:"id"`
	UserID      string    `json:"user_id" db:"user_id"`
	Name        string    `json:"name" db:"name"`
	Provider    string    `json:"provider" db:"provider"` // github, gitlab, etc.
	FullName    string    `json:"full_name" db:"full_name"` // owner/repo
	Private     bool      `json:"private" db:"private"`
	LastScanned time.Time `json:"last_scanned" db:"last_scanned"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Dependency represents a package dependency
type Dependency struct {
	ID           string    `json:"id" db:"id"`
	RepositoryID string    `json:"repository_id" db:"repository_id"`
	Name         string    `json:"name" db:"name"`
	Version      string    `json:"version" db:"version"`
	Ecosystem    string    `json:"ecosystem" db:"ecosystem"` // npm, pip, go, etc.
	Type         string    `json:"type" db:"type"` // direct, transitive
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

// Vulnerability represents a known security vulnerability
type Vulnerability struct {
	ID              string    `json:"id" db:"id"`
	CVEID           string    `json:"cve_id" db:"cve_id"`
	Title           string    `json:"title" db:"title"`
	Description     string    `json:"description" db:"description"`
	Severity        string    `json:"severity" db:"severity"` // critical, high, medium, low
	CVSSScore       float64   `json:"cvss_score" db:"cvss_score"`
	AffectedPackage string    `json:"affected_package" db:"affected_package"`
	Ecosystem       string    `json:"ecosystem" db:"ecosystem"`
	VersionRange    string    `json:"version_range" db:"version_range"` // semver range
	PatchedVersions string    `json:"patched_versions" db:"patched_versions"`
	PublishedAt     time.Time `json:"published_at" db:"published_at"`
	ExploitedAt     *time.Time `json:"exploited_at,omitempty" db:"exploited_at"` // when first exploited in wild
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// ScanResult represents the result of a repository scan
type ScanResult struct {
	ID             string    `json:"id" db:"id"`
	RepositoryID   string    `json:"repository_id" db:"repository_id"`
	DependencyID   string    `json:"dependency_id" db:"dependency_id"`
	VulnerabilityID string   `json:"vulnerability_id" db:"vulnerability_id"`
	Status         string    `json:"status" db:"status"` // vulnerable, patched, ignored
	DetectedAt     time.Time `json:"detected_at" db:"detected_at"`
	ResolvedAt     *time.Time `json:"resolved_at,omitempty" db:"resolved_at"`
}

// User represents a system user
type User struct {
	ID        string    `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Name      string    `json:"name" db:"name"`
	Plan      string    `json:"plan" db:"plan"` // free, startup, team, business
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// NotificationRule represents user notification preferences
type NotificationRule struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	Channel   string    `json:"channel" db:"channel"` // email, slack, discord, ntfy
	Enabled   bool      `json:"enabled" db:"enabled"`
	MinSeverity string  `json:"min_severity" db:"min_severity"` // only notify for this severity or higher
	Config    string    `json:"config" db:"config"` // JSON config for channel (webhook URL, etc.)
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// ScanRequest is the payload from the GitHub Action
type ScanRequest struct {
	RepoID       string       `json:"repo_id"`
	Dependencies []DependencyInfo `json:"dependencies"`
	Ecosystem    string       `json:"ecosystem"`
	Timestamp    time.Time    `json:"timestamp"`
}

// DependencyInfo is dependency information from scan
type DependencyInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"` // direct or transitive
}

// ScanResponse is returned to the GitHub Action
type ScanResponse struct {
	ScanID          string                `json:"scan_id"`
	VulnerabilitiesFound int              `json:"vulnerabilities_found"`
	Vulnerabilities []VulnerabilityMatch  `json:"vulnerabilities"`
	Timestamp       time.Time             `json:"timestamp"`
}

// VulnerabilityMatch represents a matched vulnerability
type VulnerabilityMatch struct {
	DependencyName  string  `json:"dependency_name"`
	DependencyVersion string `json:"dependency_version"`
	CVEID           string  `json:"cve_id"`
	Title           string  `json:"title"`
	Severity        string  `json:"severity"`
	CVSSScore       float64 `json:"cvss_score"`
	PatchedVersions string  `json:"patched_versions"`
	ExploitedInWild bool    `json:"exploited_in_wild"`
}
