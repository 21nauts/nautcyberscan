# GitHub Setup Guide for NautScanner

## Step 1: Create GitHub Repository

1. Go to https://github.com/new
2. Fill in:
   - **Repository name**: `NautCyberScanner`
   - **Description**: `Open Source Security Scanner for GitHub Repositories`
   - **Visibility**: Public (or Private)
   - **Do NOT initialize** with README, .gitignore, or license (we already have these)

3. Click "Create repository"

## Step 2: Push to GitHub

You'll see a page with instructions. Copy the commands or use these:

```bash
# Add GitHub as remote (replace YOUR_USERNAME with your GitHub username)
git remote add origin https://github.com/YOUR_USERNAME/NautCyberScanner.git

# Push to GitHub
git branch -M main
git push -u origin main
```

## Step 3: Set Up ngrok (Temporary API Access)

Since GitHub Actions runs in the cloud, it needs to access your API. For testing, use ngrok:

```bash
# Install ngrok
brew install ngrok

# Start ngrok (in a new terminal)
ngrok http 8081

# You'll see output like:
# Forwarding   https://abc123.ngrok.io -> http://localhost:8081
```

**Keep this terminal open!** Copy the HTTPS URL (e.g., `https://abc123.ngrok.io`)

## Step 4: Configure GitHub Secrets (Optional)

If you want to use API keys:

1. Go to your repository on GitHub
2. Click **Settings** → **Secrets and variables** → **Actions**
3. Click **New repository secret**
4. Add:
   - Name: `NAUTSCANNER_API_URL`
   - Value: `https://abc123.ngrok.io/v1` (your ngrok URL)

## Step 5: Update Workflow File

Update `.github/workflows/security-scan.yml` with your ngrok URL:

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
    name: NautScanner Security Scan

    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Run NautScanner
        uses: ./.github/actions/nautscanner
        with:
          # Replace with your ngrok URL
          api_url: https://abc123.ngrok.io/v1
          fail_on_critical: true
          min_severity: medium
```

Commit and push:

```bash
git add .github/workflows/security-scan.yml
git commit -m "Update API URL for GitHub Actions"
git push
```

## Step 6: Trigger a Scan

Make any small change and push:

```bash
echo "# Test" >> README.md
git add README.md
git commit -m "Test: Trigger security scan"
git push
```

## Step 7: View Results

1. Go to your repository on GitHub
2. Click **Actions** tab
3. You should see "Security Scan" workflow running
4. Click on the workflow to see the scan results

## Step 8: Check Dashboard

Open your local dashboard to see the scan results:

```bash
open http://localhost:3000/vulnerabilities
```

---

## Alternative: Deploy to Cloud (Production)

For production use, deploy NautScanner to a cloud provider instead of ngrok:

### Option A: Railway.app (Easiest)

1. Go to https://railway.app
2. Click "New Project" → "Deploy from GitHub"
3. Select your NautCyberScanner repository
4. Railway will auto-deploy
5. Copy the public URL and use it in your workflows

### Option B: DigitalOcean App Platform

1. Go to https://cloud.digitalocean.com
2. Create new App
3. Connect to GitHub
4. Select NautCyberScanner
5. Deploy!

### Option C: Docker on VPS

```bash
# On your VPS
git clone https://github.com/YOUR_USERNAME/NautCyberScanner.git
cd NautCyberScanner
docker-compose up -d

# Set up reverse proxy (nginx/caddy)
# Point domain to your API
```

---

## Troubleshooting

### GitHub Action Fails

**Error: "Cannot reach API"**
- Check ngrok is still running
- Verify the URL in workflow matches ngrok URL
- ngrok URLs change every time you restart - update workflow each time

**Error: "404 Not Found"**
- Make sure API URL ends with `/v1`
- Example: `https://abc123.ngrok.io/v1/scan`

### No Vulnerabilities Found

This is expected! The scanner doesn't have a vulnerability database yet.

To add real vulnerability scanning, we need to integrate with:
- OSV (Open Source Vulnerabilities)
- GitHub Advisory Database
- npm audit

Would you like me to implement this?

---

## Next Steps

1. ✅ Create GitHub repository
2. ✅ Push code
3. ✅ Set up ngrok
4. ✅ Update workflow with ngrok URL
5. ✅ Push a test commit
6. ✅ View scan results on GitHub Actions
7. ✅ **COMPLETE! GitHub Actions are now working!**
8. ⏭️ Deploy to cloud for permanent API access (optional)
9. ⏭️ Integrate real vulnerability database (to make scans meaningful)

## ✅ Setup Complete!

Your NautScanner is now successfully integrated with GitHub Actions!

**GitHub Repository**: https://github.com/21nauts/nautcyberscan
**Latest Workflow Run**: ✅ Passing
**API Status**: ✅ Running via ngrok

### What's Working:
- ✅ GitHub Actions automatically trigger on push to main/develop
- ✅ Action detects npm dependencies from package.json
- ✅ Scans are submitted to your local API via ngrok
- ✅ Results are displayed in GitHub Actions summary

### Current Limitations:
- ⚠️ Ngrok URL changes when you restart ngrok (need to update workflows each time)
- ⚠️ No real vulnerability database integrated (scans return 0 vulnerabilities)
- ⚠️ API must be running locally for scans to work

### Recommended Next Steps:
1. **For Production Use**: Deploy API to cloud (Railway, DigitalOcean, etc.)
2. **For Real Scanning**: Integrate OSV, NVD, or npm audit database
3. **For Persistence**: Ensure scan results are saved to database
