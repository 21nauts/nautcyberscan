<script lang="ts">
	import { onMount } from 'svelte';

	let vulnerabilities: any[] = [];
	let loading = true;
	let error = '';
	let filterSeverity = 'all';

	const severityColors: Record<string, string> = {
		critical: '#ff4757',
		high: '#ff6b6b',
		medium: '#ffa502',
		low: '#26de81'
	};

	onMount(async () => {
		try {
			const response = await fetch('/api/v1/vulnerabilities');
			if (response.ok) {
				vulnerabilities = await response.json();
			} else {
				error = 'Failed to load vulnerabilities';
			}
		} catch (e) {
			error = 'Error connecting to API';
		} finally {
			loading = false;
		}
	});

	$: filteredVulns = filterSeverity === 'all'
		? vulnerabilities
		: vulnerabilities.filter(v => v.severity === filterSeverity);
</script>

<svelte:head>
	<title>Vulnerabilities - NautScanner</title>
</svelte:head>

<div class="vulnerabilities">
	<div class="header">
		<h1>Vulnerabilities</h1>
		<div class="filters">
			<label>
				Severity:
				<select bind:value={filterSeverity}>
					<option value="all">All</option>
					<option value="critical">Critical</option>
					<option value="high">High</option>
					<option value="medium">Medium</option>
					<option value="low">Low</option>
				</select>
			</label>
		</div>
	</div>

	{#if loading}
		<div class="loading">Loading vulnerabilities...</div>
	{:else if error}
		<div class="error">{error}</div>
	{:else if filteredVulns.length === 0}
		<div class="empty-state">
			<div class="empty-icon">✅</div>
			<h2>No Vulnerabilities Found</h2>
			<p>
				{#if filterSeverity !== 'all'}
					No {filterSeverity} severity vulnerabilities detected.
				{:else}
					Great news! No security vulnerabilities detected in your repositories.
				{/if}
			</p>
			<div class="info-box">
				<h3>How Vulnerability Detection Works</h3>
				<ol>
					<li>NautScanner scans your dependencies on every push</li>
					<li>Compares against known vulnerability databases</li>
					<li>Alerts you immediately when issues are found</li>
					<li>Provides patch information and remediation steps</li>
				</ol>
			</div>
		</div>
	{:else}
		<div class="vuln-stats">
			<div class="stat-card glass-card critical">
				<div class="count">{vulnerabilities.filter(v => v.severity === 'critical').length}</div>
				<div class="label">Critical</div>
			</div>
			<div class="stat-card glass-card high">
				<div class="count">{vulnerabilities.filter(v => v.severity === 'high').length}</div>
				<div class="label">High</div>
			</div>
			<div class="stat-card glass-card medium">
				<div class="count">{vulnerabilities.filter(v => v.severity === 'medium').length}</div>
				<div class="label">Medium</div>
			</div>
			<div class="stat-card glass-card low">
				<div class="count">{vulnerabilities.filter(v => v.severity === 'low').length}</div>
				<div class="label">Low</div>
			</div>
		</div>

		<div class="vuln-list">
			{#each filteredVulns as vuln}
				<div class="vuln-card glass-card">
					<div class="vuln-header">
						<div>
							<h3>{vuln.title}</h3>
							<span class="cve">{vuln.cve_id}</span>
						</div>
						<span class="severity-badge" style="background: {severityColors[vuln.severity]}">
							{vuln.severity.toUpperCase()}
						</span>
					</div>
					<p class="description">{vuln.description}</p>
					<div class="vuln-meta">
						<span>Package: <strong>{vuln.affected_package}</strong></span>
						<span>CVSS: <strong>{vuln.cvss_score}</strong></span>
						<span>Published: <strong>{new Date(vuln.published_at).toLocaleDateString()}</strong></span>
					</div>
					{#if vuln.patched_versions}
						<div class="patch-info">
							✅ Patch available: {vuln.patched_versions}
						</div>
					{/if}
				</div>
			{/each}
		</div>
	{/if}
</div>

<style>
	.vulnerabilities {
		max-width: 1400px;
		margin: 0 auto;
	}

	.header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 2rem;
	}

	h1 {
		margin: 0;
		color: var(--text-primary);
		font-size: 36px;
		font-weight: 700;
	}

	.filters label {
		color: var(--text-secondary);
		font-weight: 500;
		display: flex;
		align-items: center;
		gap: 0.75rem;
	}

	.filters select {
		padding: 0.75rem 1rem;
		background: rgba(45, 49, 57, 0.6);
		border: 1px solid rgba(107, 114, 128, 0.3);
		border-radius: var(--radius-md);
		font-size: 14px;
		color: var(--text-primary);
		cursor: pointer;
		transition: all 0.3s ease;
		font-weight: 500;
	}

	.filters select:hover {
		border-color: rgba(0, 217, 255, 0.5);
	}

	.filters select:focus {
		outline: none;
		border-color: var(--cyan);
		box-shadow: 0 0 0 3px rgba(0, 217, 255, 0.1);
	}

	.loading, .error {
		text-align: center;
		padding: 3rem;
	}

	.loading {
		color: var(--text-secondary);
		font-size: 18px;
	}

	.error {
		color: var(--error);
		font-size: 18px;
		background: rgba(239, 68, 68, 0.1);
		border: 1px solid rgba(239, 68, 68, 0.3);
		border-radius: var(--radius-lg);
	}

	.empty-state {
		text-align: center;
		padding: 4rem 2rem;
	}

	.empty-icon {
		font-size: 64px;
		margin-bottom: 1.5rem;
		filter: drop-shadow(0 0 10px rgba(127, 224, 195, 0.5));
	}

	.empty-state h2 {
		margin: 0 0 1rem 0;
		color: var(--text-primary);
		font-size: 30px;
		font-weight: 600;
	}

	.empty-state p {
		color: var(--text-secondary);
		margin-bottom: 2rem;
		font-size: 16px;
	}

	.info-box {
		max-width: 600px;
		margin: 2rem auto 0;
		padding: 2rem;
		background: rgba(54, 58, 69, 0.4);
		backdrop-filter: blur(12px);
		border: 1px solid rgba(107, 114, 128, 0.2);
		border-radius: var(--radius-lg);
		text-align: left;
	}

	.info-box h3 {
		margin-top: 0;
		color: var(--text-primary);
		font-size: 20px;
		font-weight: 600;
	}

	.info-box ol {
		margin: 1rem 0;
		padding-left: 1.5rem;
		color: var(--text-secondary);
	}

	.info-box li {
		margin: 0.75rem 0;
		line-height: 1.6;
	}

	.vuln-stats {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
		gap: var(--spacing-lg);
		margin-bottom: 2rem;
	}

	.stat-card {
		padding: var(--spacing-lg);
		text-align: center;
		transition: all 0.3s ease;
	}

	.stat-card .count {
		font-size: 40px;
		font-weight: 700;
		margin-bottom: 0.5rem;
	}

	.stat-card.critical .count {
		color: var(--error);
		text-shadow: 0 0 10px rgba(239, 68, 68, 0.5);
	}
	.stat-card.high .count {
		color: #ff6b6b;
		text-shadow: 0 0 10px rgba(255, 107, 107, 0.5);
	}
	.stat-card.medium .count {
		color: var(--warning);
		text-shadow: 0 0 10px rgba(234, 179, 8, 0.5);
	}
	.stat-card.low .count {
		color: var(--success);
		text-shadow: 0 0 10px rgba(34, 197, 94, 0.5);
	}

	.stat-card .label {
		color: var(--text-secondary);
		text-transform: uppercase;
		font-size: 14px;
		font-weight: 500;
		letter-spacing: 0.5px;
	}

	.vuln-list {
		display: flex;
		flex-direction: column;
		gap: var(--spacing-lg);
	}

	.vuln-card {
		padding: var(--spacing-lg);
		transition: all 0.3s ease;
	}

	.vuln-header {
		display: flex;
		justify-content: space-between;
		align-items: start;
		margin-bottom: 1rem;
		gap: 1rem;
	}

	.vuln-card h3 {
		margin: 0 0 0.5rem 0;
		color: var(--text-primary);
		font-size: 20px;
		font-weight: 600;
	}

	.cve {
		font-size: 14px;
		color: var(--text-tertiary);
		font-family: 'Courier New', monospace;
		background: rgba(26, 29, 35, 0.6);
		padding: 2px 8px;
		border-radius: 4px;
	}

	.severity-badge {
		padding: 0.5rem 1rem;
		color: var(--text-primary);
		border-radius: var(--radius-sm);
		font-size: 12px;
		font-weight: 700;
		text-transform: uppercase;
		letter-spacing: 0.5px;
		white-space: nowrap;
		box-shadow: 0 4px 10px rgba(0, 0, 0, 0.3);
	}

	.description {
		color: var(--text-secondary);
		line-height: 1.6;
		margin: 1rem 0;
	}

	.vuln-meta {
		display: flex;
		flex-wrap: wrap;
		gap: 1.5rem;
		font-size: 14px;
		color: var(--text-secondary);
		margin: 1rem 0;
	}

	.vuln-meta strong {
		color: var(--text-primary);
		font-weight: 600;
	}

	.patch-info {
		margin-top: 1rem;
		padding: 0.75rem 1rem;
		background: rgba(34, 197, 94, 0.1);
		border-left: 4px solid var(--success);
		color: var(--success);
		border-radius: var(--radius-sm);
		font-weight: 500;
	}

	@media (max-width: 768px) {
		h1 {
			font-size: 28px;
		}

		.header {
			flex-direction: column;
			gap: 1rem;
			align-items: flex-start;
		}

		.vuln-meta {
			flex-direction: column;
			gap: 0.5rem;
		}
	}
</style>
