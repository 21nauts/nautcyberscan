package osv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/nautcyberscanner/nautscanner/pkg/models"
)

const (
	OSVAPI = "https://api.osv.dev/v1"
)

// Client handles communication with OSV API
type Client struct {
	httpClient *http.Client
}

// New creates a new OSV API client
func New() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// QueryRequest represents an OSV API query request
type QueryRequest struct {
	Package Package `json:"package"`
	Version string  `json:"version,omitempty"`
}

// Package represents package information
type Package struct {
	Name      string `json:"name"`
	Ecosystem string `json:"ecosystem"`
}

// QueryResponse represents an OSV API query response
type QueryResponse struct {
	Vulns []Vulnerability `json:"vulns"`
}

// Vulnerability represents an OSV vulnerability
type Vulnerability struct {
	ID       string   `json:"id"`
	Summary  string   `json:"summary"`
	Details  string   `json:"details"`
	Severity []Severity `json:"severity,omitempty"`
	Affected []Affected `json:"affected"`
	References []Reference `json:"references,omitempty"`
}

// Severity represents CVSS severity information
type Severity struct {
	Type  string `json:"type"`
	Score string `json:"score"`
}

// Affected represents affected package versions
type Affected struct {
	Package           Package `json:"package"`
	Ranges            []Range `json:"ranges"`
	Versions          []string `json:"versions,omitempty"`
	DatabaseSpecific  map[string]interface{} `json:"database_specific,omitempty"`
	EcosystemSpecific map[string]interface{} `json:"ecosystem_specific,omitempty"`
}

// Range represents version ranges
type Range struct {
	Type   string  `json:"type"`
	Events []Event `json:"events"`
}

// Event represents a range event (introduced/fixed)
type Event struct {
	Introduced string `json:"introduced,omitempty"`
	Fixed      string `json:"fixed,omitempty"`
}

// Reference represents a vulnerability reference
type Reference struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// QueryVulnerabilities queries OSV API for vulnerabilities in a specific package version
func (c *Client) QueryVulnerabilities(packageName, version, ecosystem string) ([]models.VulnerabilityMatch, error) {
	// Map ecosystem names to OSV format
	osvEcosystem := mapEcosystem(ecosystem)

	reqBody := QueryRequest{
		Package: Package{
			Name:      packageName,
			Ecosystem: osvEcosystem,
		},
		Version: version,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", OSVAPI+"/query", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to query OSV API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OSV API returned status %d", resp.StatusCode)
	}

	var queryResp QueryResponse
	if err := json.NewDecoder(resp.Body).Decode(&queryResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Convert OSV vulnerabilities to our model
	var matches []models.VulnerabilityMatch
	for _, vuln := range queryResp.Vulns {
		match := models.VulnerabilityMatch{
			CVEID:             vuln.ID,
			Title:             vuln.Summary,
			DependencyName:    packageName,
			DependencyVersion: version,
			Severity:          extractSeverity(vuln.Severity),
			CVSSScore:         extractCVSSScore(vuln.Severity),
			PatchedVersions:   extractPatchedVersions(vuln.Affected),
			ExploitedInWild:   false, // OSV doesn't provide this
		}
		matches = append(matches, match)
	}

	return matches, nil
}

// BatchQuery queries multiple packages at once
func (c *Client) BatchQuery(dependencies []models.DependencyInfo, ecosystem string) ([]models.VulnerabilityMatch, error) {
	var allMatches []models.VulnerabilityMatch

	for _, dep := range dependencies {
		// Clean version string (remove ^, ~, etc.)
		version := cleanVersion(dep.Version)

		matches, err := c.QueryVulnerabilities(dep.Name, version, ecosystem)
		if err != nil {
			// Log error but continue with other dependencies
			fmt.Printf("Warning: failed to query %s: %v\n", dep.Name, err)
			continue
		}
		allMatches = append(allMatches, matches...)
	}

	return allMatches, nil
}

// Helper functions

func mapEcosystem(ecosystem string) string {
	ecosystemMap := map[string]string{
		"npm":    "npm",
		"python": "PyPI",
		"go":     "Go",
		"rust":   "crates.io",
		"ruby":   "RubyGems",
	}

	if mapped, ok := ecosystemMap[ecosystem]; ok {
		return mapped
	}
	return ecosystem
}

func cleanVersion(version string) string {
	// Remove common version prefixes
	version = trimPrefix(version, "^")
	version = trimPrefix(version, "~")
	version = trimPrefix(version, ">=")
	version = trimPrefix(version, "<=")
	version = trimPrefix(version, ">")
	version = trimPrefix(version, "<")
	version = trimPrefix(version, "=")
	return version
}

func trimPrefix(s, prefix string) string {
	if len(s) > 0 && s[0:1] == prefix {
		return s[1:]
	}
	return s
}

func extractSeverity(severities []Severity) string {
	if len(severities) == 0 {
		return "UNKNOWN"
	}

	// Prefer CVSS scores
	for _, sev := range severities {
		if sev.Type == "CVSS_V3" || sev.Type == "CVSS_V2" {
			score := extractCVSSScore(severities)
			if score >= 9.0 {
				return "CRITICAL"
			} else if score >= 7.0 {
				return "HIGH"
			} else if score >= 4.0 {
				return "MEDIUM"
			} else if score > 0 {
				return "LOW"
			}
		}
	}

	return "UNKNOWN"
}

func extractCVSSScore(severities []Severity) float64 {
	if len(severities) == 0 {
		return 0.0
	}

	// Parse CVSS score from severity
	for _, sev := range severities {
		if sev.Type == "CVSS_V3" || sev.Type == "CVSS_V2" {
			// CVSS score format is usually "CVSS:3.1/AV:N/AC:L/..."
			// We need to extract the base score
			// For simplicity, return 0.0 for now
			// TODO: Implement proper CVSS parsing
			return 0.0
		}
	}

	return 0.0
}

func extractPatchedVersions(affected []Affected) string {
	var patchedVersions []string

	for _, aff := range affected {
		for _, r := range aff.Ranges {
			for _, event := range r.Events {
				if event.Fixed != "" {
					patchedVersions = append(patchedVersions, event.Fixed)
				}
			}
		}
	}

	if len(patchedVersions) == 0 {
		return "None"
	}

	// Join versions with comma
	result := ""
	for i, v := range patchedVersions {
		if i > 0 {
			result += ", "
		}
		result += v
	}
	return result
}
