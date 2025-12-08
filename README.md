# NautScanner

🛡️ **Open Source Security Scanner for Your Repositories**

NautScanner automatically scans your repositories for dependency vulnerabilities and tracks library versions. Stay ahead of security threats like React2Shell (CVE-2025-55182) with automated monitoring and instant notifications.

## Features

- 🔍 **Automated Dependency Scanning** - Scan npm, pip, Go modules, and more
- 🚨 **Real-time Vulnerability Detection** - Get alerted when new CVEs affect your code
- ⚡ **Fast Exploitation Alerts** - Know when vulnerabilities are being exploited in the wild
- 📊 **Beautiful Dashboard** - Visualize your security posture across all repos
- 🔔 **Multi-channel Notifications** - Email, Slack, Discord, or self-hosted Ntfy
- 🐳 **Self-hosted or SaaS** - Run locally with Docker or use our hosted service
- 🎯 **GitHub Actions Integration** - Automatic scanning on every push

## Quick Start

### Option 1: Docker (Recommended for Self-Hosting)

```bash
# Clone the repository
git clone https://github.com/nautcyberscanner/nautscanner.git
cd nautscanner

# Copy environment file
cp .env.example .env

# Start all services
docker-compose up -d

# Visit http://localhost:3000
```

### Option 2: Local Development

**Prerequisites:**
- Go 1.21+
- Node.js 20+
- PostgreSQL 16+

```bash
# Install dependencies
make install-deps

# Start database
docker-compose up -d postgres

# Run backend
cd backend
go run cmd/server/main.go

# In another terminal, run frontend
cd frontend
npm run dev
```

Visit:
- Frontend: http://localhost:3000
- API: http://localhost:8080

## GitHub Action Setup

Add NautScanner to your repository to automatically scan on every push:

### 1. Create workflow file

Create `.github/workflows/nautscanner.yml`:

```yaml
name: Security Scan

on: [push, pull_request]

jobs:
  scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: nautcyberscanner/scan-action@v1
        with:
          api_key: ${{ secrets.NAUTSCANNER_API_KEY }}
          fail_on_critical: true
```

### 2. Add API Key (Optional for Self-Hosted)

For the SaaS version, add your API key to repository secrets:

1. Go to Settings → Secrets → Actions
2. Add `NAUTSCANNER_API_KEY` with your key from https://nautscanner.io

For self-hosted installations, point to your own API:

```yaml
- uses: nautcyberscanner/scan-action@v1
  with:
    api_url: https://your-nautscanner-instance.com/v1
```

## Architecture

```
┌─────────────────────────────────────┐
│   GitHub Repository                 │
│   ┌──────────────────────┐          │
│   │ GitHub Action        │          │
│   │ (Extracts deps only) │          │
│   └──────────┬───────────┘          │
└──────────────┼──────────────────────┘
               │ POST /api/v1/scan
               │ { dependencies: [...] }
               ▼
┌──────────────────────────────────────┐
│   NautScanner Backend (Go)           │
│   ┌────────────────────┐             │
│   │ Vulnerability      │             │
│   │ Matcher Engine     │             │
│   └────────┬───────────┘             │
│            │                         │
│   ┌────────▼───────────┐             │
│   │ PostgreSQL         │             │
│   │ (Vulns + Scans)    │             │
│   └────────────────────┘             │
└──────────────┬───────────────────────┘
               │
               ▼
┌──────────────────────────────────────┐
│   Frontend (SvelteKit)               │
│   - Dashboard                        │
│   - Vulnerability Reports            │
│   - Settings                         │
└──────────────────────────────────────┘
```

## Project Structure

```
nautscanner/
├── backend/                 # Go API server
│   ├── cmd/server/         # Main application
│   ├── internal/           # Internal packages
│   │   ├── api/           # HTTP handlers
│   │   ├── database/      # Database layer
│   │   ├── matcher/       # Vulnerability matching
│   │   └── notifier/      # Notification system
│   └── pkg/               # Public packages
│       ├── models/        # Data models
│       └── config/        # Configuration
├── frontend/               # SvelteKit web app
│   └── src/
│       ├── routes/        # Pages
│       └── lib/           # Components
├── .github/
│   ├── workflows/         # CI/CD
│   └── actions/           # GitHub Action
├── docs/                  # Documentation
└── docker-compose.yml     # Local development
```

## Development Commands

```bash
make help          # Show all commands
make dev           # Start development environment
make build         # Build backend and frontend
make test          # Run tests
make docker-up     # Start Docker containers
make docker-down   # Stop Docker containers
make clean         # Clean build artifacts
```

## Configuration

Key environment variables (see `.env.example`):

```bash
# Database
DATABASE_URL=postgres://nautscanner:nautscanner@localhost:5432/nautscanner

# Server
SERVER_PORT=8080
ENVIRONMENT=development

# Notifications
SMTP_HOST=smtp.example.com
SLACK_WEBHOOK=https://hooks.slack.com/...
NTFY_URL=https://ntfy.sh
```

## Supported Ecosystems

Current support:
- ✅ JavaScript/TypeScript (npm, yarn, pnpm)
- ✅ Python (pip, poetry)
- ✅ Go (go.mod)

Coming soon:
- 🔜 Java/Kotlin (Maven, Gradle)
- 🔜 Ruby (Bundler)
- 🔜 Rust (Cargo)
- 🔜 PHP (Composer)

## Vulnerability Data Sources

NautScanner aggregates data from:
- GitHub Security Advisories
- National Vulnerability Database (NVD)
- npm audit
- Python Safety DB
- Go Vulnerability Database

## Pricing

### Open Source (Free Forever)
- Self-hosted Docker deployment
- Unlimited repositories
- All core features
- Community support

### SaaS Plans

**Startup ($9/month)**
- Up to 10 repositories
- Daily auto-scans
- Email + Slack notifications
- 7-day history

**Team ($29/month)**
- Up to 50 repositories
- Hourly auto-scans
- All notification channels
- 90-day history
- Priority support

**Business ($99/month)**
- Unlimited repositories
- Real-time scanning
- API access
- 1-year audit trail
- SLA support

## Contributing

We welcome contributions! See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## Security

Found a security issue? Email security@nautscanner.io (GPG key available)

## License

MIT License - see [LICENSE](LICENSE)

## Inspiration

This project was inspired by the rapid exploitation of React2Shell (CVE-2025-55182), where threat actors began exploiting the vulnerability within hours of public disclosure. Read more in the [AWS Security Blog](https://aws.amazon.com/blogs/security/china-nexus-cyber-threat-groups-rapidly-exploit-react2shell-vulnerability-cve-2025-55182/).

## Links

- 🌐 Website: https://nautscanner.io
- 📚 Documentation: https://docs.nautscanner.io
- 💬 Discord: https://discord.gg/nautscanner
- 🐦 Twitter: [@nautscanner](https://twitter.com/nautscanner)

---

Made with ❤️ by the NautScanner team
