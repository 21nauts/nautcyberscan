<script lang="ts">
	import '../app.css';
</script>

<div class="app">
	<header class="header">
		<nav class="nav-container">
			<div class="logo">
				<h1>NautScanner</h1>
				<span class="version">v1.0</span>
			</div>
			<ul class="nav-menu">
				<li><a href="/" class="nav-link">Dashboard</a></li>
				<li><a href="/repositories" class="nav-link">Repositories</a></li>
				<li><a href="/vulnerabilities" class="nav-link">Vulnerabilities</a></li>
				<li><a href="/settings" class="nav-link">Settings</a></li>
			</ul>
		</nav>
	</header>

	<main class="main-content">
		<slot />
	</main>

	<footer class="footer">
		<p>&copy; 2024 NautScanner - Open Source Security Scanner</p>
	</footer>
</div>

<style>
	.app {
		display: flex;
		flex-direction: column;
		min-height: 100vh;
	}

	.header {
		background: rgba(45, 49, 57, 0.4);
		backdrop-filter: blur(12px);
		-webkit-backdrop-filter: blur(12px);
		border-bottom: 1px solid rgba(107, 114, 128, 0.2);
		padding: 1rem 0;
		position: sticky;
		top: 0;
		z-index: 100;
	}

	.nav-container {
		max-width: 1280px;
		margin: 0 auto;
		padding: 0 2rem;
		display: flex;
		justify-content: space-between;
		align-items: center;
	}

	.logo {
		display: flex;
		align-items: center;
		gap: 0.75rem;
	}

	.logo h1 {
		color: var(--cyan);
		font-size: 24px;
		font-weight: 700;
		margin: 0;
		text-shadow: 0 0 20px rgba(0, 217, 255, 0.5);
	}

	.version {
		padding: 4px 8px;
		background: rgba(0, 217, 255, 0.2);
		border: 1px solid rgba(0, 217, 255, 0.3);
		border-radius: var(--radius-sm);
		font-size: 12px;
		font-weight: 500;
		color: var(--cyan);
	}

	.nav-menu {
		list-style: none;
		display: flex;
		gap: 0.5rem;
		margin: 0;
		padding: 0;
	}

	.nav-link {
		display: block;
		padding: 0.75rem 1.25rem;
		color: var(--text-secondary);
		text-decoration: none;
		border-radius: var(--radius-md);
		font-weight: 500;
		font-size: 14px;
		transition: all 0.2s ease;
		position: relative;
	}

	.nav-link:hover {
		color: var(--cyan);
		background: rgba(0, 217, 255, 0.1);
	}

	.nav-link:global(.active) {
		color: var(--text-primary);
		background: rgba(0, 217, 255, 0.2);
	}

	.main-content {
		flex: 1;
		max-width: 1280px;
		width: 100%;
		margin: 0 auto;
		padding: 2rem;
	}

	.footer {
		background: rgba(45, 49, 57, 0.4);
		backdrop-filter: blur(12px);
		border-top: 1px solid rgba(107, 114, 128, 0.2);
		color: var(--text-tertiary);
		text-align: center;
		padding: 1.5rem;
		margin-top: auto;
	}

	.footer p {
		margin: 0;
		font-size: 14px;
		color: var(--text-tertiary);
	}

	@media (max-width: 768px) {
		.nav-container {
			flex-direction: column;
			gap: 1rem;
		}

		.nav-menu {
			width: 100%;
			justify-content: center;
		}

		.nav-link {
			padding: 0.5rem 1rem;
			font-size: 13px;
		}

		.main-content {
			padding: 1rem;
		}
	}
</style>
