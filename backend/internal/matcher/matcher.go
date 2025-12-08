package matcher

import (
	"database/sql"
	"fmt"

	"github.com/nautcyberscanner/nautscanner/pkg/models"
)

// Matcher handles vulnerability matching logic
type Matcher struct {
	db *sql.DB
}

// New creates a new vulnerability matcher
func New(db *sql.DB) *Matcher {
	return &Matcher{db: db}
}

// MatchDependencies finds vulnerabilities for given dependencies
func (m *Matcher) MatchDependencies(dependencies []models.DependencyInfo, ecosystem string) ([]models.VulnerabilityMatch, error) {
	var matches []models.VulnerabilityMatch

	for _, dep := range dependencies {
		vulns, err := m.findVulnerabilities(dep.Name, dep.Version, ecosystem)
		if err != nil {
			return nil, fmt.Errorf("failed to find vulnerabilities for %s: %w", dep.Name, err)
		}
		matches = append(matches, vulns...)
	}

	return matches, nil
}

// findVulnerabilities queries the database for matching vulnerabilities
func (m *Matcher) findVulnerabilities(packageName, version, ecosystem string) ([]models.VulnerabilityMatch, error) {
	query := `
		SELECT
			cve_id,
			title,
			severity,
			cvss_score,
			patched_versions,
			exploited_at IS NOT NULL as exploited_in_wild
		FROM vulnerabilities
		WHERE affected_package = $1
			AND ecosystem = $2
			AND $3 ~ version_range
	`

	rows, err := m.db.Query(query, packageName, ecosystem, version)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []models.VulnerabilityMatch
	for rows.Next() {
		var match models.VulnerabilityMatch
		match.DependencyName = packageName
		match.DependencyVersion = version

		err := rows.Scan(
			&match.CVEID,
			&match.Title,
			&match.Severity,
			&match.CVSSScore,
			&match.PatchedVersions,
			&match.ExploitedInWild,
		)
		if err != nil {
			return nil, err
		}

		matches = append(matches, match)
	}

	return matches, nil
}

// IsVersionVulnerable checks if a specific version is vulnerable
// This is a placeholder - implement proper semver range checking
func (m *Matcher) IsVersionVulnerable(version, versionRange string) bool {
	// TODO: Implement proper semver range matching
	// For now, just return false
	return false
}
