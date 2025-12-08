<script lang="ts">
	import { onMount } from 'svelte';

	let stats = {
		totalRepos: 0,
		vulnerabilities: 0,
		critical: 0,
		lastScan: 'Never'
	};

	onMount(async () => {
		// TODO: Fetch actual stats from API
		stats = {
			totalRepos: 0,
			vulnerabilities: 0,
			critical: 0,
			lastScan: 'Never'
		};
	});
</script>

<svelte:head>
	<title>Dashboard - NautScanner</title>
</svelte:head>

<div class="dashboard">
	<div class="hero">
		<h1>Security Dashboard</h1>
		<p class="hero-subtitle">Real-time vulnerability monitoring for your repositories</p>
	</div>

	<div class="stats-grid">
		<div class="stat-card glass-card">
			<div class="stat-icon">📦</div>
			<div class="stat-content">
				<h3 class="stat-label">Repositories</h3>
				<p class="stat-value">{stats.totalRepos}</p>
			</div>
		</div>

		<div class="stat-card glass-card">
			<div class="stat-icon">⚠️</div>
			<div class="stat-content">
				<h3 class="stat-label">Vulnerabilities</h3>
				<p class="stat-value">{stats.vulnerabilities}</p>
			</div>
		</div>

		<div class="stat-card glass-card critical">
			<div class="stat-icon">🔴</div>
			<div class="stat-content">
				<h3 class="stat-label">Critical</h3>
				<p class="stat-value">{stats.critical}</p>
			</div>
		</div>

		<div class="stat-card glass-card">
			<div class="stat-icon">🔍</div>
			<div class="stat-content">
				<h3 class="stat-label">Last Scan</h3>
				<p class="stat-value small">{stats.lastScan}</p>
			</div>
		</div>
	</div>

	<div class="quick-start">
		<h2>Get Started</h2>
		<div class="steps">
			<div class="step glass-card">
				<div class="step-number">1</div>
				<h3>Install GitHub Action</h3>
				<p>Add NautScanner to your repository workflows</p>
				<a href="/docs/install" class="btn-primary">View Instructions</a>
			</div>
			<div class="step glass-card">
				<div class="step-number">2</div>
				<h3>Configure Notifications</h3>
				<p>Set up alerts for security vulnerabilities</p>
				<a href="/settings" class="btn-primary">Configure</a>
			</div>
			<div class="step glass-card">
				<div class="step-number">3</div>
				<h3>Monitor Your Repos</h3>
				<p>Track dependencies and vulnerabilities</p>
				<a href="/repositories" class="btn-primary">View Repositories</a>
			</div>
		</div>
	</div>
</div>

<style>
	.dashboard {
		padding: 0;
	}

	.hero {
		text-align: center;
		margin-bottom: 3rem;
	}

	.hero h1 {
		margin: 0 0 0.75rem 0;
		color: var(--text-primary);
		font-size: 36px;
		font-weight: 700;
	}

	.hero-subtitle {
		color: var(--text-secondary);
		font-size: 18px;
		margin: 0;
	}

	.stats-grid {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
		gap: var(--spacing-lg);
		margin-bottom: 3rem;
	}

	.stat-card {
		padding: var(--spacing-lg);
		display: flex;
		align-items: center;
		gap: var(--spacing-md);
		transition: all 0.3s ease;
	}

	.stat-card.critical {
		border-color: rgba(239, 68, 68, 0.3);
	}

	.stat-icon {
		font-size: 40px;
		opacity: 0.9;
	}

	.stat-content {
		flex: 1;
	}

	.stat-label {
		margin: 0 0 0.5rem 0;
		color: var(--text-secondary);
		font-size: 14px;
		font-weight: 500;
		text-transform: uppercase;
		letter-spacing: 0.5px;
	}

	.stat-value {
		font-size: 36px;
		font-weight: 700;
		margin: 0;
		color: var(--cyan);
		text-shadow: 0 0 10px rgba(0, 217, 255, 0.3);
	}

	.stat-value.small {
		font-size: 18px;
		color: var(--text-primary);
		text-shadow: none;
	}

	.quick-start {
		margin-top: 3rem;
	}

	.quick-start h2 {
		margin: 0 0 2rem 0;
		color: var(--text-primary);
		text-align: center;
		font-size: 30px;
		font-weight: 600;
	}

	.steps {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
		gap: var(--spacing-lg);
	}

	.step {
		padding: var(--spacing-xl);
		text-align: center;
		transition: all 0.3s ease;
	}

	.step-number {
		width: 60px;
		height: 60px;
		margin: 0 auto 1.5rem;
		background: linear-gradient(135deg, var(--cyan), var(--cyan-dark));
		color: var(--text-primary);
		border-radius: 50%;
		display: flex;
		align-items: center;
		justify-content: center;
		font-size: 24px;
		font-weight: 700;
		box-shadow: 0 0 20px rgba(0, 217, 255, 0.4);
	}

	.step h3 {
		margin: 0 0 0.75rem 0;
		color: var(--text-primary);
		font-size: 20px;
		font-weight: 600;
	}

	.step p {
		color: var(--text-secondary);
		margin: 0 0 1.5rem 0;
		line-height: 1.6;
	}

	.btn-primary {
		display: inline-block;
		padding: 0.75rem 1.5rem;
		background: linear-gradient(135deg, var(--cyan), var(--cyan-dark));
		color: var(--text-primary);
		text-decoration: none;
		border-radius: var(--radius-md);
		transition: all 0.3s ease;
		font-weight: 600;
		box-shadow: 0 4px 15px rgba(0, 217, 255, 0.3);
	}

	.btn-primary:hover {
		transform: translateY(-2px);
		box-shadow: 0 6px 20px rgba(0, 217, 255, 0.5);
	}

	@media (max-width: 768px) {
		.hero h1 {
			font-size: 28px;
		}

		.stat-card {
			flex-direction: column;
			text-align: center;
		}
	}
</style>
