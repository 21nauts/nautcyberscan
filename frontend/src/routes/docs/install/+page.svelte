<svelte:head>
	<title>Installation Guide - NautScanner</title>
</svelte:head>

<div class="docs">
	<div class="sidebar">
		<h3>Documentation</h3>
		<nav>
			<a href="/docs/install" class="active">Installation</a>
			<a href="/docs/quickstart">Quick Start</a>
			<a href="/docs/api">API Reference</a>
			<a href="/docs/faq">FAQ</a>
		</nav>
	</div>

	<div class="content glass-card">
		<h1>Installation Guide</h1>

		<section>
			<h2>GitHub Action Installation</h2>
			<p>Add NautScanner to your repository to automatically scan for vulnerabilities on every push.</p>

			<h3>Step 1: Create Workflow File</h3>
			<p>Create a new file in your repository:</p>
			<pre><code>.github/workflows/security-scan.yml</code></pre>

			<h3>Step 2: Add Configuration</h3>
			<p>Copy this workflow configuration:</p>
			<pre><code>{`name: Security Scan

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]
  schedule:
    # Run daily at 2 AM UTC
    - cron: '0 2 * * *'

jobs:
  scan:
    runs-on: ubuntu-latest
    name: NautScanner Security Scan

    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Run NautScanner
        uses: nautcyberscanner/scan-action@v1
        with:
          api_key: \${{ secrets.NAUTSCANNER_API_KEY }}
          fail_on_critical: true
          min_severity: medium`}</code></pre>

			<h3>Step 3: Configure API Key (Optional)</h3>
			<p>For self-hosted installations, point to your API:</p>
			<pre><code>{`      - name: Run NautScanner
        uses: nautcyberscanner/scan-action@v1
        with:
          api_url: http://localhost:8081/v1
          # No API key needed for self-hosted`}</code></pre>

			<h3>Step 4: Commit and Push</h3>
			<pre><code>{`git add .github/workflows/security-scan.yml
git commit -m "Add NautScanner security scanning"
git push`}</code></pre>

			<p>That's it! Your repository will now be scanned automatically.</p>
		</section>

		<section>
			<h2>Self-Hosted Installation</h2>

			<h3>Docker (Recommended)</h3>
			<pre><code>{`# Clone repository
git clone https://github.com/nautcyberscanner/nautscanner.git
cd nautscanner

# Start services
docker-compose up -d

# Access dashboard
open http://localhost:3000`}</code></pre>

			<h3>Manual Installation</h3>
			<pre><code>{`# Prerequisites: Go 1.21+, Node.js 20+, PostgreSQL 16+

# Install dependencies
make install-deps

# Start database
docker-compose up -d postgres

# Run backend
cd backend && go run cmd/server/main.go

# Run frontend (in another terminal)
cd frontend && npm run dev`}</code></pre>
		</section>

		<section>
			<h2>Configuration Options</h2>

			<table>
				<thead>
					<tr>
						<th>Option</th>
						<th>Description</th>
						<th>Default</th>
					</tr>
				</thead>
				<tbody>
					<tr>
						<td><code>api_key</code></td>
						<td>Your NautScanner API key</td>
						<td>-</td>
					</tr>
					<tr>
						<td><code>api_url</code></td>
						<td>API endpoint URL</td>
						<td>https://api.nautscanner.io/v1</td>
					</tr>
					<tr>
						<td><code>fail_on_critical</code></td>
						<td>Fail build on critical vulns</td>
						<td>false</td>
					</tr>
					<tr>
						<td><code>min_severity</code></td>
						<td>Minimum severity to report</td>
						<td>medium</td>
					</tr>
				</tbody>
			</table>
		</section>

		<section>
			<h2>Next Steps</h2>
			<ul>
				<li><a href="/docs/quickstart">Quick Start Guide</a> - Get up and running in 5 minutes</li>
				<li><a href="/settings">Configure Notifications</a> - Set up Slack, email, or Discord alerts</li>
				<li><a href="/repositories">View Repositories</a> - Monitor your scanned repositories</li>
				<li><a href="/docs/api">API Documentation</a> - Integrate with your tools</li>
			</ul>
		</section>
	</div>
</div>

<style>
	.docs {
		display: grid;
		grid-template-columns: 250px 1fr;
		gap: 3rem;
		max-width: 1400px;
		margin: 0 auto;
	}

	.sidebar {
		position: sticky;
		top: 2rem;
		height: fit-content;
		padding: var(--spacing-lg);
		background: rgba(45, 49, 57, 0.4);
		backdrop-filter: blur(12px);
		border: 1px solid rgba(107, 114, 128, 0.2);
		border-radius: var(--radius-lg);
	}

	.sidebar h3 {
		margin: 0 0 1rem 0;
		color: var(--text-primary);
		font-size: 18px;
		font-weight: 600;
	}

	.sidebar nav {
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
	}

	.sidebar a {
		padding: 0.75rem 1rem;
		color: var(--text-secondary);
		text-decoration: none;
		border-radius: var(--radius-md);
		transition: all 0.3s ease;
		font-size: 14px;
		font-weight: 500;
	}

	.sidebar a:hover {
		background: rgba(0, 217, 255, 0.1);
		color: var(--cyan);
	}

	.sidebar a.active {
		background: linear-gradient(135deg, var(--cyan), var(--cyan-dark));
		color: var(--text-primary);
		box-shadow: 0 4px 10px rgba(0, 217, 255, 0.3);
	}

	.content {
		padding: 3rem;
	}

	h1 {
		margin: 0 0 2rem 0;
		color: var(--text-primary);
		font-size: 36px;
		font-weight: 700;
	}

	h2 {
		margin: 2.5rem 0 1rem 0;
		color: var(--text-primary);
		border-bottom: 2px solid var(--cyan);
		padding-bottom: 0.75rem;
		font-size: 28px;
		font-weight: 600;
	}

	h3 {
		margin: 1.5rem 0 0.75rem 0;
		color: var(--text-primary);
		font-size: 20px;
		font-weight: 600;
	}

	section {
		margin: 2.5rem 0;
	}

	p {
		line-height: 1.6;
		color: var(--text-secondary);
		margin: 1rem 0;
		font-size: 15px;
	}

	pre {
		background: rgba(26, 29, 35, 0.8);
		border: 1px solid rgba(107, 114, 128, 0.3);
		color: var(--text-primary);
		padding: 1.5rem;
		border-radius: var(--radius-md);
		overflow-x: auto;
		margin: 1rem 0;
		box-shadow: inset 0 2px 8px rgba(0, 0, 0, 0.3);
	}

	code {
		font-family: 'Courier New', monospace;
		font-size: 14px;
	}

	table {
		width: 100%;
		border-collapse: collapse;
		margin: 1.5rem 0;
		background: rgba(45, 49, 57, 0.4);
		backdrop-filter: blur(12px);
		border: 1px solid rgba(107, 114, 128, 0.2);
		border-radius: var(--radius-md);
		overflow: hidden;
	}

	th, td {
		padding: 1rem;
		text-align: left;
		border-bottom: 1px solid rgba(107, 114, 128, 0.2);
	}

	th {
		background: rgba(26, 29, 35, 0.6);
		font-weight: 600;
		color: var(--text-primary);
		font-size: 14px;
		text-transform: uppercase;
		letter-spacing: 0.5px;
	}

	td {
		color: var(--text-secondary);
		font-size: 14px;
	}

	td code {
		background: rgba(0, 217, 255, 0.1);
		padding: 0.25rem 0.5rem;
		border-radius: var(--radius-sm);
		color: var(--cyan);
		border: 1px solid rgba(0, 217, 255, 0.2);
	}

	ul {
		line-height: 2;
		margin: 1rem 0;
		color: var(--text-secondary);
	}

	ul a {
		color: var(--cyan);
		text-decoration: none;
		transition: all 0.2s ease;
	}

	ul a:hover {
		text-decoration: underline;
		color: var(--cyan-light);
	}

	@media (max-width: 768px) {
		.docs {
			grid-template-columns: 1fr;
			gap: 2rem;
		}

		.sidebar {
			position: static;
		}

		h1 {
			font-size: 28px;
		}

		h2 {
			font-size: 24px;
		}

		.content {
			padding: 2rem;
		}
	}
</style>
