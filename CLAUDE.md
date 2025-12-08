# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview
NautCyberScanner is a vulnerability scanning and library version tracking tool for software repositories. It provides both an open-source self-hosted version and a commercial SaaS offering.

### Core Functionality
1. **Repository Integration**: Connect to GitHub repos and scan dependencies
2. **Version Tracking**: Identify all library versions (React, Next.js, etc.) across projects
3. **Feature Discovery**: Fetch and display major features for each library version
4. **Vulnerability Monitoring**: Web scraping or subscription-based vulnerability tracking
5. **Automated Reporting**: Generate reports on library versions and security vulnerabilities
6. **Real-time Alerts**: Notify users when their libraries have known vulnerabilities

### Deployment Models
- **Open Source**: Dockerized application for local deployment (community-supported)
- **SaaS**: Commercial hosted version with automatic scanning and premium features

## Project Blueprint
The initial project requirements are documented in:
- `docs/2025-12-08_blueprint_nautscanner.md` - Original project goals and requirements

### Key Reference
The project was inspired by the React2Shell vulnerability (CVE-2025-55182):
https://aws.amazon.com/blogs/security/china-nexus-cyber-threat-groups-rapidly-exploit-react2shell-vulnerability-cve-2025-55182/

## Architecture Considerations

### Backend Requirements
- **Repository Scanning**: Need to parse package.json, requirements.txt, go.mod, etc.
- **Vulnerability Database**: Integration with CVE databases, GitHub Security Advisories, npm audit, etc.
- **Web Scraping**: Automated collection of vulnerability announcements
- **API Integrations**: GitHub API for repository access, library registries for version info
- **Database**: Store repository metadata, scan results, vulnerability data, user preferences

### Frontend Requirements
- **Dashboard**: Visual representation of scanned repositories and their health
- **Library Details**: Interactive views of library versions and available features
- **Vulnerability Alerts**: Clear presentation of security issues with severity levels
- **Report Generation**: PDF/HTML reports for compliance and auditing

### Docker Deployment
- Self-contained application with all dependencies
- Configuration via environment variables
- Volume mounts for persistent data
- Optional integration with external vulnerability feeds

## Development Phases (Suggested)

### Phase 1: Core Scanning
- Repository connection (GitHub initially)
- Dependency file parsing (package.json, requirements.txt, etc.)
- Version detection and tracking
- Basic UI for repository list and versions

### Phase 2: Vulnerability Integration
- CVE database integration
- GitHub Security Advisories
- npm/PyPI security feeds
- Vulnerability matching against detected versions

### Phase 3: Feature Discovery
- Library changelog scraping/API integration
- Feature extraction from release notes
- UI for displaying major features per version

### Phase 4: Automation & Reporting
- Scheduled scanning
- Alert system
- Report generation
- Docker packaging

### Phase 5: SaaS Features
- User authentication and multi-tenancy
- Payment integration
- Auto-scanning webhooks
- Advanced reporting and analytics

## Data Sources to Consider

### Vulnerability Databases
- National Vulnerability Database (NVD)
- GitHub Security Advisories
- npm audit / yarn audit
- Snyk Vulnerability Database
- OSV (Open Source Vulnerabilities)

### Library Information
- npm registry API
- PyPI API
- Maven Central
- RubyGems API
- Go pkg.go.dev
- GitHub releases and changelogs

## Security Considerations
- Secure storage of GitHub tokens and API keys
- Rate limiting for external API calls
- Sandboxed execution for repository analysis
- No execution of untrusted code from scanned repositories
- Secure handling of vulnerability data
- GDPR compliance for SaaS version (user data)

## Testing Strategy
- Unit tests for parsing different dependency file formats
- Integration tests for API connections
- Mock data for vulnerability matching
- End-to-end tests for scanning workflows
- Load testing for SaaS deployment

## Current Status
**MVP Scaffold Complete** - Full project structure implemented with:
- Go backend with API server, database, and routing
- SvelteKit frontend with dashboard UI
- PostgreSQL database with migrations
- Docker Compose development environment
- GitHub Action for automated scanning
- Complete documentation

## Build & Run Commands

### Local Development
```bash
# Install dependencies
make install-deps

# Start development environment
make dev

# Build production binaries
make build

# Run tests
make test
```

### Docker
```bash
# Start all services
docker-compose up -d

# Stop services
docker-compose down

# View logs
docker-compose logs -f
```

### Individual Services
```bash
# Backend only
cd backend && go run cmd/server/main.go

# Frontend only
cd frontend && npm run dev
```

## Key Files

### Backend Entry Point
- `backend/cmd/server/main.go` - Application entry point
- `backend/internal/api/routes.go` - API route definitions
- `backend/internal/database/db.go` - Database migrations

### Frontend Entry Points
- `frontend/src/routes/+page.svelte` - Dashboard page
- `frontend/src/routes/+layout.svelte` - App layout
- `frontend/vite.config.ts` - API proxy configuration

### Configuration
- `.env.example` - Environment variable template
- `docker-compose.yml` - Local development stack
- `Makefile` - Build automation

## API Endpoints

All endpoints are prefixed with `/api/v1`:

- `POST /scan` - Submit scan from GitHub Action
- `GET /repositories` - List repositories
- `GET /repositories/{id}` - Get repository details
- `GET /repositories/{id}/vulnerabilities` - Get vulnerabilities
- `GET /vulnerabilities` - List all vulnerabilities
- `GET /vulnerabilities/{id}` - Get vulnerability details
- `GET /users/me` - Get current user
- `PUT /users/me/notifications` - Update notification settings

## Database Schema

Tables: users, repositories, dependencies, vulnerabilities, scan_results, notification_rules

All tables use UUIDs as primary keys and include created_at/updated_at timestamps.

## Next Implementation Tasks

1. **Vulnerability Matching** - Implement semver range checking in `backend/internal/matcher/matcher.go`
2. **Data Sources** - Integrate GitHub Security Advisories, npm audit, etc.
3. **Authentication** - Add JWT-based auth for API endpoints
4. **Frontend Pages** - Build repositories, vulnerabilities, and settings pages
5. **Notifications** - Implement email, Slack, Discord, and Ntfy notifications
6. **GitHub Integration** - Complete the GitHub App setup for automatic repository discovery
