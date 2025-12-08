package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// NewConnection creates a new database connection
func NewConnection(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	return db, nil
}

// RunMigrations executes all database migrations
func RunMigrations(db *sql.DB) error {
	migrations := []string{
		createUsersTable,
		createRepositoriesTable,
		createDependenciesTable,
		createVulnerabilitiesTable,
		createScanResultsTable,
		createNotificationRulesTable,
		createIndexes,
	}

	for i, migration := range migrations {
		if _, err := db.Exec(migration); err != nil {
			return fmt.Errorf("failed to execute migration %d: %w", i+1, err)
		}
	}

	return nil
}

const createUsersTable = `
CREATE TABLE IF NOT EXISTS users (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	email VARCHAR(255) UNIQUE NOT NULL,
	name VARCHAR(255) NOT NULL,
	plan VARCHAR(50) DEFAULT 'free',
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const createRepositoriesTable = `
CREATE TABLE IF NOT EXISTS repositories (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	name VARCHAR(255) NOT NULL,
	provider VARCHAR(50) DEFAULT 'github',
	full_name VARCHAR(512) NOT NULL,
	private BOOLEAN DEFAULT false,
	last_scanned TIMESTAMP,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	UNIQUE(user_id, full_name)
);
`

const createDependenciesTable = `
CREATE TABLE IF NOT EXISTS dependencies (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	repository_id UUID NOT NULL REFERENCES repositories(id) ON DELETE CASCADE,
	name VARCHAR(255) NOT NULL,
	version VARCHAR(100) NOT NULL,
	ecosystem VARCHAR(50) NOT NULL,
	type VARCHAR(50) DEFAULT 'direct',
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	UNIQUE(repository_id, name, ecosystem)
);
`

const createVulnerabilitiesTable = `
CREATE TABLE IF NOT EXISTS vulnerabilities (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	cve_id VARCHAR(50) UNIQUE NOT NULL,
	title VARCHAR(500) NOT NULL,
	description TEXT,
	severity VARCHAR(20) NOT NULL,
	cvss_score DECIMAL(3,1),
	affected_package VARCHAR(255) NOT NULL,
	ecosystem VARCHAR(50) NOT NULL,
	version_range VARCHAR(255),
	patched_versions VARCHAR(255),
	published_at TIMESTAMP NOT NULL,
	exploited_at TIMESTAMP,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const createScanResultsTable = `
CREATE TABLE IF NOT EXISTS scan_results (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	repository_id UUID NOT NULL REFERENCES repositories(id) ON DELETE CASCADE,
	dependency_id UUID NOT NULL REFERENCES dependencies(id) ON DELETE CASCADE,
	vulnerability_id UUID NOT NULL REFERENCES vulnerabilities(id) ON DELETE CASCADE,
	status VARCHAR(50) DEFAULT 'vulnerable',
	detected_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	resolved_at TIMESTAMP,
	UNIQUE(repository_id, dependency_id, vulnerability_id)
);
`

const createNotificationRulesTable = `
CREATE TABLE IF NOT EXISTS notification_rules (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	channel VARCHAR(50) NOT NULL,
	enabled BOOLEAN DEFAULT true,
	min_severity VARCHAR(20) DEFAULT 'medium',
	config JSONB,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const createIndexes = `
CREATE INDEX IF NOT EXISTS idx_repositories_user_id ON repositories(user_id);
CREATE INDEX IF NOT EXISTS idx_dependencies_repository_id ON dependencies(repository_id);
CREATE INDEX IF NOT EXISTS idx_dependencies_name ON dependencies(name);
CREATE INDEX IF NOT EXISTS idx_vulnerabilities_cve_id ON vulnerabilities(cve_id);
CREATE INDEX IF NOT EXISTS idx_vulnerabilities_package ON vulnerabilities(affected_package, ecosystem);
CREATE INDEX IF NOT EXISTS idx_scan_results_repository_id ON scan_results(repository_id);
CREATE INDEX IF NOT EXISTS idx_scan_results_status ON scan_results(status);
CREATE INDEX IF NOT EXISTS idx_notification_rules_user_id ON notification_rules(user_id);
`
