package api

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/nautcyberscanner/nautscanner/pkg/config"
)

// SetupRoutes configures all API routes
func SetupRoutes(router *mux.Router, db *sql.DB, cfg *config.Config) {
	// Health check
	router.HandleFunc("/health", HealthCheck).Methods("GET")

	// API v1
	api := router.PathPrefix("/api/v1").Subrouter()

	// Scan endpoint (called by GitHub Actions)
	scanHandler := NewScanHandler(db, cfg)
	api.HandleFunc("/scan", scanHandler.SubmitScan).Methods("POST")

	// Repository endpoints
	repoHandler := NewRepositoryHandler(db)
	api.HandleFunc("/repositories", repoHandler.ListRepositories).Methods("GET")
	api.HandleFunc("/repositories/{id}", repoHandler.GetRepository).Methods("GET")
	api.HandleFunc("/repositories/{id}/vulnerabilities", repoHandler.GetRepositoryVulnerabilities).Methods("GET")

	// Vulnerability endpoints
	vulnHandler := NewVulnerabilityHandler(db)
	api.HandleFunc("/vulnerabilities", vulnHandler.ListVulnerabilities).Methods("GET")
	api.HandleFunc("/vulnerabilities/{id}", vulnHandler.GetVulnerability).Methods("GET")

	// User endpoints
	userHandler := NewUserHandler(db)
	api.HandleFunc("/users/me", userHandler.GetCurrentUser).Methods("GET")
	api.HandleFunc("/users/me/notifications", userHandler.UpdateNotificationSettings).Methods("PUT")

	// CORS middleware
	router.Use(corsMiddleware)
}
