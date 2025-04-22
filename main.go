// Entry point for the API service
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/yuvakkrishnan/user-activity-logger/api"
	"github.com/yuvakkrishnan/user-activity-logger/internal/cache"
	"github.com/yuvakkrishnan/user-activity-logger/internal/db"
	"github.com/yuvakkrishnan/user-activity-logger/internal/logger"
)

// HealthCheckHandler handles the health check endpoint
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize logger
	logger.InitLogger("info")
	log.Println("Initializing API server...")

	// Initialize Redis and PostgreSQL
	cache.InitRedis()
	db.InitDB()

	// Set up router
	r := mux.NewRouter()
	r.Use(api.LoggingMiddleware)

	// Define routes
	r.HandleFunc("/health", HealthCheckHandler).Methods("GET")
	apiRoutes := r.PathPrefix("/api").Subrouter()
	apiRoutes.Use(api.JWTMiddleware)
	apiRoutes.HandleFunc("/activity", api.UserActivityHandler).Methods("POST")

	// Start the server
	port := ":8080"
	log.Printf("🚀 API listening on %s", port)
	fmt.Printf("Starting the server on port %s...\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
