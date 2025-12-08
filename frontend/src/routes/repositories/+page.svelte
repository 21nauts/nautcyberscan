<script lang="ts">
	import { onMount } from 'svelte';

	let repositories: any[] = [];
	let loading = true;
	let error = '';

	onMount(async () => {
		try {
			const response = await fetch('/api/v1/repositories');
			if (response.ok) {
				repositories = await response.json();
			} else {
				error = 'Failed to load repositories';
			}
		} catch (e) {
			error = 'Error connecting to API';
		} finally {
			loading = false;
		}
	});
</script>

<svelte:head>
	<title>Repositories - NautScanner</title>
</svelte:head>

<div class="repositories">
	<div class="header">
		<h1>Repositories</h1>
		<button class="btn-primary">+ Add Repository</button>
	</div>

	{#if loading}
		<div class="loading">Loading repositories...</div>
	{:else if error}
		<div class="error">{error}</div>
	{:else if repositories.length === 0}
		<div class="empty-state">
			<div class="empty-icon">📦</div>
			<h2>No Repositories Yet</h2>
			<p>Get started by adding your first repository to scan for vulnerabilities.</p>

			<div class="getting-started">
				<h3>Quick Setup</h3>
				<ol>
					<li>Add the NautScanner GitHub Action to your repository</li>
					<li>Push code to trigger the first scan</li>
					<li>Your repository will appear here automatically</li>
				</ol>
				<a href="/docs/install" class="btn">View Installation Guide</a>
			</div>
		</div>
	{:else}
		<div class="repo-grid">
			{#each repositories as repo}
				<div class="repo-card glass-card">
					<div class="repo-header">
						<h3>{repo.name}</h3>
						<span class="badge">{repo.private ? 'Private' : 'Public'}</span>
					</div>
					<p class="repo-path">{repo.full_name}</p>
					<div class="repo-stats">
						<div class="stat">
							<span class="label">Last Scan:</span>
							<span class="value">{repo.last_scanned || 'Never'}</span>
						</div>
						<div class="stat">
							<span class="label">Vulnerabilities:</span>
							<span class="value critical">0</span>
						</div>
					</div>
					<button class="btn-secondary">View Details</button>
				</div>
			{/each}
		</div>
	{/if}
</div>

<style>
	.repositories {
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

	.btn-primary {
		background: linear-gradient(135deg, var(--cyan), var(--cyan-dark));
		color: var(--text-primary);
		border: none;
		padding: 0.75rem 1.5rem;
		border-radius: var(--radius-md);
		font-size: 1rem;
		font-weight: 600;
		cursor: pointer;
		transition: all 0.3s ease;
		box-shadow: 0 4px 15px rgba(0, 217, 255, 0.3);
	}

	.btn-primary:hover {
		transform: translateY(-2px);
		box-shadow: 0 6px 20px rgba(0, 217, 255, 0.5);
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
		opacity: 0.8;
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

	.getting-started {
		max-width: 600px;
		margin: 2rem auto 0;
		padding: 2rem;
		background: rgba(54, 58, 69, 0.4);
		backdrop-filter: blur(12px);
		border: 1px solid rgba(107, 114, 128, 0.2);
		border-radius: var(--radius-lg);
		text-align: left;
	}

	.getting-started h3 {
		margin-top: 0;
		color: var(--text-primary);
		font-size: 20px;
		font-weight: 600;
	}

	.getting-started ol {
		margin: 1rem 0;
		padding-left: 1.5rem;
		color: var(--text-secondary);
	}

	.getting-started li {
		margin: 0.75rem 0;
		line-height: 1.6;
	}

	.btn {
		display: inline-block;
		margin-top: 1rem;
		padding: 0.75rem 1.5rem;
		background: linear-gradient(135deg, var(--cyan), var(--cyan-dark));
		color: var(--text-primary);
		text-decoration: none;
		border-radius: var(--radius-md);
		transition: all 0.3s ease;
		font-weight: 600;
		box-shadow: 0 4px 15px rgba(0, 217, 255, 0.3);
	}

	.btn:hover {
		transform: translateY(-2px);
		box-shadow: 0 6px 20px rgba(0, 217, 255, 0.5);
	}

	.repo-grid {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
		gap: var(--spacing-lg);
	}

	.repo-card {
		padding: var(--spacing-lg);
		transition: all 0.3s ease;
	}

	.repo-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 0.75rem;
	}

	.repo-card h3 {
		margin: 0;
		color: var(--text-primary);
		font-size: 20px;
		font-weight: 600;
	}

	.badge {
		padding: 0.25rem 0.75rem;
		background: rgba(127, 224, 195, 0.2);
		border: 1px solid rgba(127, 224, 195, 0.3);
		border-radius: 12px;
		font-size: 12px;
		font-weight: 500;
		color: var(--mint);
	}

	.repo-path {
		color: var(--text-tertiary);
		font-size: 14px;
		margin: 0.5rem 0 1rem 0;
		font-family: 'Courier New', monospace;
	}

	.repo-stats {
		margin: 1rem 0;
		padding: 1rem;
		background: rgba(26, 29, 35, 0.4);
		border-radius: var(--radius-md);
	}

	.stat {
		display: flex;
		justify-content: space-between;
		margin: 0.5rem 0;
		font-size: 14px;
	}

	.stat .label {
		color: var(--text-secondary);
	}

	.stat .value {
		font-weight: 600;
		color: var(--text-primary);
	}

	.stat .value.critical {
		color: var(--error);
		text-shadow: 0 0 8px rgba(239, 68, 68, 0.4);
	}

	.btn-secondary {
		width: 100%;
		padding: 0.75rem;
		background: transparent;
		border: 2px solid var(--cyan);
		color: var(--cyan);
		border-radius: var(--radius-md);
		cursor: pointer;
		transition: all 0.3s ease;
		margin-top: 1rem;
		font-weight: 600;
	}

	.btn-secondary:hover {
		background: var(--cyan);
		color: var(--bg-primary);
		box-shadow: 0 4px 15px rgba(0, 217, 255, 0.3);
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
	}
</style>
