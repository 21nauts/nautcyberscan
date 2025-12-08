<script lang="ts">
	let notifications = {
		email: true,
		slack: false,
		discord: false,
		ntfy: false
	};

	let emailAddress = '';
	let slackWebhook = '';
	let discordWebhook = '';
	let ntfyTopic = '';

	let minSeverity = 'medium';
	let saved = false;

	function saveSettings() {
		// TODO: Save to API
		saved = true;
		setTimeout(() => saved = false, 3000);
	}
</script>

<svelte:head>
	<title>Settings - NautScanner</title>
</svelte:head>

<div class="settings">
	<h1>Settings</h1>

	<div class="settings-grid">
		<!-- Notification Settings -->
		<section class="settings-section glass-card">
			<h2>🔔 Notification Channels</h2>
			<p class="description">Choose how you want to receive security alerts</p>

			<div class="setting-item">
				<label class="toggle">
					<input type="checkbox" bind:checked={notifications.email} />
					<span class="slider"></span>
					<span class="label-text">Email Notifications</span>
				</label>
				{#if notifications.email}
					<input
						type="email"
						bind:value={emailAddress}
						placeholder="your@email.com"
						class="input-field"
					/>
				{/if}
			</div>

			<div class="setting-item">
				<label class="toggle">
					<input type="checkbox" bind:checked={notifications.slack} />
					<span class="slider"></span>
					<span class="label-text">Slack</span>
				</label>
				{#if notifications.slack}
					<input
						type="text"
						bind:value={slackWebhook}
						placeholder="https://hooks.slack.com/..."
						class="input-field"
					/>
				{/if}
			</div>

			<div class="setting-item">
				<label class="toggle">
					<input type="checkbox" bind:checked={notifications.discord} />
					<span class="slider"></span>
					<span class="label-text">Discord</span>
				</label>
				{#if notifications.discord}
					<input
						type="text"
						bind:value={discordWebhook}
						placeholder="https://discord.com/api/webhooks/..."
						class="input-field"
					/>
				{/if}
			</div>

			<div class="setting-item">
				<label class="toggle">
					<input type="checkbox" bind:checked={notifications.ntfy} />
					<span class="slider"></span>
					<span class="label-text">Ntfy (Self-hosted)</span>
				</label>
				{#if notifications.ntfy}
					<input
						type="text"
						bind:value={ntfyTopic}
						placeholder="my-topic-name"
						class="input-field"
					/>
				{/if}
			</div>
		</section>

		<!-- Alert Settings -->
		<section class="settings-section glass-card">
			<h2>⚠️ Alert Settings</h2>
			<p class="description">Configure when you want to be notified</p>

			<div class="setting-item">
				<label for="severity">Minimum Severity</label>
				<select id="severity" bind:value={minSeverity} class="select-field">
					<option value="low">Low (All vulnerabilities)</option>
					<option value="medium">Medium (Medium and above)</option>
					<option value="high">High (High and critical only)</option>
					<option value="critical">Critical (Critical only)</option>
				</select>
				<p class="help-text">Only receive alerts for vulnerabilities at or above this severity level</p>
			</div>
		</section>

		<!-- API Settings -->
		<section class="settings-section glass-card">
			<h2>🔑 API Settings</h2>
			<p class="description">Manage your API access</p>

			<div class="setting-item">
				<label>API Key</label>
				<div class="api-key-display">
					<code>naut_********************************</code>
					<button class="btn-copy">Copy</button>
				</div>
				<p class="help-text">Use this key to authenticate API requests</p>
			</div>

			<div class="setting-item">
				<button class="btn-danger">Regenerate API Key</button>
				<p class="help-text warning">⚠️ This will invalidate your current key</p>
			</div>
		</section>

		<!-- Account Settings -->
		<section class="settings-section glass-card">
			<h2>👤 Account</h2>
			<p class="description">Manage your account settings</p>

			<div class="setting-item">
				<label>Current Plan</label>
				<div class="plan-info">
					<strong>Free (Open Source)</strong>
					<p>Self-hosted, unlimited repositories</p>
				</div>
			</div>

			<div class="setting-item">
				<label>Usage</label>
				<div class="usage-stats">
					<div class="usage-item">
						<span>Repositories:</span>
						<strong>0 / Unlimited</strong>
					</div>
					<div class="usage-item">
						<span>Scans this month:</span>
						<strong>0</strong>
					</div>
				</div>
			</div>
		</section>
	</div>

	<!-- Save Button -->
	<div class="actions">
		<button class="btn-primary" on:click={saveSettings}>
			{saved ? '✅ Saved!' : 'Save Settings'}
		</button>
	</div>
</div>

<style>
	.settings {
		max-width: 1200px;
		margin: 0 auto;
	}

	h1 {
		margin: 0 0 2rem 0;
		color: var(--text-primary);
		font-size: 36px;
		font-weight: 700;
	}

	.settings-grid {
		display: grid;
		gap: var(--spacing-xl);
	}

	.settings-section {
		padding: var(--spacing-xl);
	}

	.settings-section h2 {
		margin: 0 0 0.75rem 0;
		color: var(--text-primary);
		font-size: 24px;
		font-weight: 600;
	}

	.description {
		color: var(--text-secondary);
		margin: 0 0 1.5rem 0;
		font-size: 14px;
	}

	.setting-item {
		margin: 1.5rem 0;
	}

	.setting-item label {
		display: block;
		margin-bottom: 0.5rem;
		color: var(--text-primary);
		font-weight: 500;
	}

	/* Toggle Switch */
	.toggle {
		display: flex;
		align-items: center;
		cursor: pointer;
		user-select: none;
	}

	.toggle input {
		display: none;
	}

	.slider {
		position: relative;
		width: 50px;
		height: 26px;
		background: rgba(107, 114, 128, 0.4);
		border-radius: 26px;
		transition: 0.3s;
		margin-right: 1rem;
		border: 1px solid rgba(107, 114, 128, 0.3);
	}

	.slider::before {
		content: '';
		position: absolute;
		width: 20px;
		height: 20px;
		left: 3px;
		bottom: 2px;
		background: var(--text-tertiary);
		border-radius: 50%;
		transition: 0.3s;
	}

	.toggle input:checked + .slider {
		background: linear-gradient(135deg, var(--cyan), var(--cyan-dark));
		border-color: var(--cyan);
		box-shadow: 0 0 10px rgba(0, 217, 255, 0.3);
	}

	.toggle input:checked + .slider::before {
		transform: translateX(22px);
		background: var(--text-primary);
	}

	.label-text {
		font-weight: 500;
		color: var(--text-primary);
	}

	/* Input Fields */
	.input-field, .select-field {
		width: 100%;
		padding: 0.75rem;
		background: rgba(26, 29, 35, 0.6);
		border: 1px solid rgba(107, 114, 128, 0.3);
		border-radius: var(--radius-md);
		font-size: 14px;
		color: var(--text-primary);
		margin-top: 0.5rem;
		transition: all 0.3s ease;
	}

	.input-field:focus, .select-field:focus {
		outline: none;
		border-color: var(--cyan);
		box-shadow: 0 0 0 3px rgba(0, 217, 255, 0.1);
	}

	.help-text {
		font-size: 13px;
		color: var(--text-tertiary);
		margin: 0.5rem 0 0 0;
		line-height: 1.5;
	}

	.help-text.warning {
		color: var(--error);
	}

	/* API Key */
	.api-key-display {
		display: flex;
		gap: 1rem;
		align-items: center;
		padding: 0.75rem 1rem;
		background: rgba(26, 29, 35, 0.6);
		border: 1px solid rgba(107, 114, 128, 0.3);
		border-radius: var(--radius-md);
		margin-top: 0.5rem;
	}

	.api-key-display code {
		flex: 1;
		color: var(--cyan);
		font-family: 'Courier New', monospace;
		font-size: 14px;
	}

	.btn-copy {
		padding: 0.5rem 1rem;
		background: linear-gradient(135deg, var(--cyan), var(--cyan-dark));
		color: var(--text-primary);
		border: none;
		border-radius: var(--radius-sm);
		cursor: pointer;
		font-weight: 600;
		transition: all 0.3s ease;
		box-shadow: 0 4px 10px rgba(0, 217, 255, 0.2);
	}

	.btn-copy:hover {
		transform: translateY(-1px);
		box-shadow: 0 6px 15px rgba(0, 217, 255, 0.3);
	}

	/* Plan Info */
	.plan-info {
		padding: 1.25rem;
		background: rgba(34, 197, 94, 0.1);
		border-left: 4px solid var(--success);
		border-radius: var(--radius-md);
	}

	.plan-info strong {
		color: var(--success);
		font-size: 18px;
	}

	.plan-info p {
		margin: 0.5rem 0 0 0;
		color: var(--text-secondary);
	}

	/* Usage Stats */
	.usage-stats {
		padding: 1.25rem;
		background: rgba(26, 29, 35, 0.6);
		border: 1px solid rgba(107, 114, 128, 0.3);
		border-radius: var(--radius-md);
	}

	.usage-item {
		display: flex;
		justify-content: space-between;
		margin: 0.75rem 0;
	}

	.usage-item span {
		color: var(--text-secondary);
	}

	.usage-item strong {
		color: var(--text-primary);
		font-weight: 600;
	}

	/* Buttons */
	.actions {
		margin-top: 2rem;
		display: flex;
		justify-content: flex-end;
	}

	.btn-primary {
		padding: 0.75rem 2rem;
		background: linear-gradient(135deg, var(--cyan), var(--cyan-dark));
		color: var(--text-primary);
		border: none;
		border-radius: var(--radius-md);
		font-size: 16px;
		font-weight: 600;
		cursor: pointer;
		transition: all 0.3s ease;
		box-shadow: 0 4px 15px rgba(0, 217, 255, 0.3);
	}

	.btn-primary:hover {
		transform: translateY(-2px);
		box-shadow: 0 6px 20px rgba(0, 217, 255, 0.5);
	}

	.btn-danger {
		padding: 0.5rem 1rem;
		background: var(--error);
		color: var(--text-primary);
		border: none;
		border-radius: var(--radius-md);
		cursor: pointer;
		font-weight: 600;
		transition: all 0.3s ease;
		box-shadow: 0 4px 10px rgba(239, 68, 68, 0.2);
	}

	.btn-danger:hover {
		transform: translateY(-1px);
		box-shadow: 0 6px 15px rgba(239, 68, 68, 0.4);
	}

	@media (max-width: 768px) {
		h1 {
			font-size: 28px;
		}
	}
</style>
