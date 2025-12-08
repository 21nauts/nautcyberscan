#!/bin/bash

# Test script to scan the NautCyberScanner repository itself

echo "🔍 Scanning NautCyberScanner repository..."
echo ""

# Prepare scan request
cat > /tmp/scan-request.json <<EOF
{
  "repository": {
    "name": "NautCyberScanner",
    "full_name": "user/NautCyberScanner",
    "url": "file:///Volumes/DevHub_ext/factory/02-development/NautCyberScanner"
  },
  "dependencies": [
    {
      "name": "@sveltejs/adapter-auto",
      "version": "^3.0.0",
      "ecosystem": "npm"
    },
    {
      "name": "@sveltejs/kit",
      "version": "^2.0.0",
      "ecosystem": "npm"
    },
    {
      "name": "@sveltejs/vite-plugin-svelte",
      "version": "^3.0.0",
      "ecosystem": "npm"
    },
    {
      "name": "svelte",
      "version": "^4.2.7",
      "ecosystem": "npm"
    },
    {
      "name": "svelte-check",
      "version": "^3.6.0",
      "ecosystem": "npm"
    },
    {
      "name": "typescript",
      "version": "^5.0.0",
      "ecosystem": "npm"
    },
    {
      "name": "vite",
      "version": "^5.0.3",
      "ecosystem": "npm"
    },
    {
      "name": "@sveltejs/adapter-node",
      "version": "^2.0.0",
      "ecosystem": "npm"
    }
  ]
}
EOF

echo "📦 Dependencies to scan:"
echo "  - @sveltejs/kit ^2.0.0"
echo "  - svelte ^4.2.7"
echo "  - vite ^5.0.3"
echo "  - typescript ^5.0.0"
echo "  - ... and 4 more"
echo ""

echo "🚀 Sending scan request to API..."
response=$(curl -s -X POST http://localhost:8081/api/v1/scan \
  -H "Content-Type: application/json" \
  -d @/tmp/scan-request.json)

echo ""
echo "📊 Scan Response:"
echo "$response" | jq '.' 2>/dev/null || echo "$response"

# Clean up
rm /tmp/scan-request.json

echo ""
echo "✅ Scan complete!"
echo "📱 View results at: http://localhost:3000/vulnerabilities"
