# NautScanner Usage Guide

## How to Use NautScanner

There are **3 ways** to scan repositories with NautScanner:

---

## 1️⃣ Local Manual Scan (Testing)

**Best for:** Testing the scanner locally on any folder

### Quick Test
```bash
# Run the test script (scans NautCyberScanner itself)
./test-scan.sh
```

### Scan Any Project
Create a scan request for any project by calling the API directly:

```bash
curl -X POST http://localhost:8081/api/v1/scan \
  -H "Content-Type: application/json" \
  -d '{
    "repository": {
      "name": "my-project",
      "full_name": "user/my-project",
      "url": "file:///path/to/project"
    },
    "dependencies": [
      {
        "name": "react",
        "version": "18.2.0",
        "ecosystem": "npm"
      }
    ]
  }'
```

**View results:** http://localhost:3000/vulnerabilities

---

## 2️⃣ GitHub Action (Automatic)

**Best for:** Continuous scanning of GitHub repositories

### Step 1: Add Workflow File

In **any GitHub repository** you want to scan, create:

`.github/workflows/security-scan.yml`

```yaml
name: Security Scan

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Run NautScanner
        uses: nautcyberscanner/scan-action@v1
        with:
          api_url: https://your-nautscanner-api.com/v1
          fail_on_critical: true
          min_severity: medium
```

### Step 2: Configure API URL

**Option A: Use Self-Hosted (Current Setup)**

You need to expose your local API to the internet:

```bash
# Using ngrok (install first: brew install ngrok)
ngrok http 8081

# Copy the HTTPS URL and use it in the workflow:
# api_url: https://abc123.ngrok.io/v1
```

**Option B: Deploy to Cloud**

Deploy NautScanner to a cloud provider:
- DigitalOcean
- AWS
- Heroku
- Render.com

Then use that URL in your workflow.

### Step 3: Push Code

```bash
git add .github/workflows/security-scan.yml
git commit -m "Add NautScanner security scanning"
git push
```

The scan runs automatically on every push!

---

## 3️⃣ Using NautScanner as a Composite Action (Local Repos)

**Best for:** Scanning local repositories with GitHub Actions workflow

This is what you have in `.github/actions/nautscanner/action.yml`

### Setup for Local Testing

1. **Create a workflow in the repo you want to scan:**

```yaml
# In any repo: .github/workflows/scan.yml
name: Local Security Scan

on:
  push:
    branches: [ main ]

jobs:
  scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      # Copy NautScanner action to this repo
      - name: Copy NautScanner action
        run: |
          mkdir -p .github/actions/nautscanner
          # Copy the action files from your NautScanner repo

      - name: Run scan
        uses: ./.github/actions/nautscanner
        with:
          api_url: http://localhost:8081/v1
```

2. **Run locally with act:**

```bash
# Install act (GitHub Actions locally)
brew install act

# Run the workflow
act -j scan
```

---

## 🎯 Recommended Setup for You

Since you want to **scan repositories without Git**, here's what I recommend:

### Option A: Web Interface (I'll build this)

I can create a simple web form where you:
1. Enter a folder path
2. Click "Scan"
3. View results immediately

### Option B: Command Line Tool

```bash
# I can create a CLI tool
nautscanner scan /path/to/project

# Or scan specific files
nautscanner scan . --include package.json
```

### Option C: Watch Folder

```bash
# Auto-scan when files change
nautscanner watch /path/to/project
```

---

## Current Limitations

⚠️ **The scanner currently doesn't check actual vulnerabilities** because:
- No vulnerability database is connected yet
- Need to integrate with OSV, NVD, or npm audit

Would you like me to:
1. ✅ **Add a web UI for scanning local folders** (easiest for you)
2. ✅ **Integrate real vulnerability database** (makes it actually useful)
3. ✅ **Create a CLI tool** (scan from terminal)

Which would you prefer?
