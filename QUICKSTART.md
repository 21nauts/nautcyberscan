# NautScanner - Quick Start Guide

Get up and running with NautScanner in 5 minutes!

## Prerequisites

Choose one option:

### Option A: Docker (Easiest)
- Docker Desktop installed
- That's it!

### Option B: Local Development
- Go 1.21+
- Node.js 20+
- PostgreSQL 16+

## 5-Minute Setup

### Step 1: Clone and Setup

```bash
# Clone the repository
git clone https://github.com/nautcyberscanner/nautscanner.git
cd nautscanner

# Copy environment file
cp .env.example .env
```

### Step 2: Start Services

**Docker (Recommended):**
```bash
docker-compose up -d
```

**Local Development:**
```bash
# Install dependencies
make install-deps

# Start database only
docker-compose up -d postgres

# Start backend (in one terminal)
cd backend
go run cmd/server/main.go

# Start frontend (in another terminal)
cd frontend
npm run dev
```

### Step 3: Access the Application

Open your browser:
- **Frontend**: http://localhost:3000
- **API**: http://localhost:8080/health

You should see:
- The NautScanner dashboard at localhost:3000
- API health check returning `{"status":"healthy"}`

## Test the GitHub Action

### Step 1: Add to Your Repository

Create `.github/workflows/security-scan.yml` in any of your repos:

```yaml
name: Security Scan

on: [push]

jobs:
  scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: nautcyberscanner/scan-action@v1
        with:
          api_url: http://localhost:8080/v1  # Or your deployed URL
```

### Step 2: Push and Watch

```bash
git add .github/workflows/security-scan.yml
git commit -m "Add NautScanner"
git push
```

Check the Actions tab to see the scan results!

## Common Commands

```bash
# View logs
docker-compose logs -f

# Stop everything
docker-compose down

# Restart after code changes
docker-compose restart backend
docker-compose restart frontend

# Run tests
make test

# Build production binaries
make build
```

## Troubleshooting

### Port Already in Use

If ports 3000, 8080, or 5432 are in use:

```bash
# Edit docker-compose.yml and change ports
# For example:
ports:
  - "3001:3000"  # Change frontend to 3001
```

### Database Connection Failed

```bash
# Check if PostgreSQL is running
docker-compose ps

# View database logs
docker-compose logs postgres

# Restart database
docker-compose restart postgres
```

### Frontend Not Loading

```bash
# Check if backend is running
curl http://localhost:8080/health

# View frontend logs
docker-compose logs frontend

# Rebuild frontend
cd frontend && npm run build
```

## Next Steps

1. **Explore the Dashboard** - Navigate to http://localhost:3000
2. **Read the Docs** - Check out README.md for detailed information
3. **Add Repositories** - Connect your first repository
4. **Configure Notifications** - Set up Slack, email, or Ntfy alerts
5. **Customize** - Modify the code to fit your needs!

## Quick Architecture Overview

```
┌─────────────────┐
│  Your Repo      │
│  (GitHub Action)│
└────────┬────────┘
         │ Sends dependency info
         ▼
┌─────────────────┐
│  Backend (Go)   │  ←→  PostgreSQL
│  Port 8080      │
└────────┬────────┘
         │ JSON API
         ▼
┌─────────────────┐
│  Frontend       │
│  (SvelteKit)    │
│  Port 3000      │
└─────────────────┘
```

## Need Help?

- 📚 **Documentation**: See README.md
- 🐛 **Issues**: https://github.com/nautcyberscanner/nautscanner/issues
- 💬 **Discord**: https://discord.gg/nautscanner
- 📧 **Email**: support@nautscanner.io

Happy scanning! 🛡️
