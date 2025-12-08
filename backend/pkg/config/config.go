package config

import (
	"fmt"
	"os"
)

// Config holds all configuration for the application
type Config struct {
	ServerPort      string
	DatabaseURL     string
	Environment     string
	JWTSecret       string

	// GitHub Integration
	GitHubAppID     string
	GitHubAppSecret string

	// Notification Services
	SMTPHost        string
	SMTPPort        string
	SMTPUser        string
	SMTPPassword    string
	SlackWebhook    string
	NtfyURL         string
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	cfg := &Config{
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://nautscanner:nautscanner@localhost:5432/nautscanner?sslmode=disable"),
		Environment: getEnv("ENVIRONMENT", "development"),
		JWTSecret:   getEnv("JWT_SECRET", "change-me-in-production"),

		GitHubAppID:     getEnv("GITHUB_APP_ID", ""),
		GitHubAppSecret: getEnv("GITHUB_APP_SECRET", ""),

		SMTPHost:     getEnv("SMTP_HOST", ""),
		SMTPPort:     getEnv("SMTP_PORT", "587"),
		SMTPUser:     getEnv("SMTP_USER", ""),
		SMTPPassword: getEnv("SMTP_PASSWORD", ""),
		SlackWebhook: getEnv("SLACK_WEBHOOK", ""),
		NtfyURL:      getEnv("NTFY_URL", "https://ntfy.sh"),
	}

	// Validate required fields in production
	if cfg.Environment == "production" {
		if cfg.JWTSecret == "change-me-in-production" {
			return nil, fmt.Errorf("JWT_SECRET must be set in production")
		}
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
