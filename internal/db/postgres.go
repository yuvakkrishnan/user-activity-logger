// PostgreSQL integration
package db

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"

	"github.com/yuvakkrishnan/user-activity-logger/internal/logger"
	"github.com/yuvakkrishnan/user-activity-logger/pkg/models"

	_ "github.com/lib/pq" // ✅ needed to register "postgres" driver
)

// Global variable to hold the database connection
var DB *sql.DB

// InitDB initializes the PostgreSQL database connection
func InitDB() {
	// Fetch connection details from environment variables (set in Docker Compose)

	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		port = "5432"
	}
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		user = "postgres"
	}
	password := url.QueryEscape(os.Getenv("POSTGRES_PASSWORD")) // Apple#Tree -> Apple%23Tree
	//password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")


	// Validate required environment variables
	if host == "" || port == "" || user == "" || password == "" || dbName == "" {

		logger.Log.Fatalf("❌ Missing required environment variables for PostgreSQL connection")
	}

	// Build the connection string
	//connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbName)

	fmt.Println("🔗 DSN:", connStr)

	// // Open the database connection working old code...
	// var err error
	// DB, err = sql.Open("postgres", connStr)
	// if err != nil {
	// 	logger.Log.Fatalf("❌ Postgres open failed: %v", err)
	// }

	// Test the connection

	// New DB connection
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Log.Fatalf("❌ Postgres open failed: %v", err)
	}
	fmt.Println("✅ sql.Open done")

	// Ping
	fmt.Println("➡️ Pinging DB...")
	err = DB.Ping()
	if err != nil {
		fmt.Printf("❌ Postgres ping failed: %v\n", err) // Add this line
		logger.Log.Errorf("❌ Postgres ping failed: %v", err)
		if closeErr := DB.Close(); closeErr != nil {
			logger.Log.Errorf("❌ Failed to close DB connection: %v", closeErr)
		}
		os.Exit(1)
	}

	logger.Log.Info("✅ Connected to Postgres")
	//fmt.Println("=========================", "Code entered by After validate user", "=========================")

}

// CloseDB closes the PostgreSQL database connection
// SaveActivity saves a user activity to the database
func SaveActivity(a models.UserActivity) error {
	query := `
		INSERT INTO activity_log (user_id, action, metadata, timestamp)
		VALUES ($1, $2, $3, $4)
	`
	_, err := DB.Exec(query, a.UserID, a.Action, a.Metadata, a.TimeStamp)
	if err != nil {
		logger.Log.Errorf("❌ Failed to save activity: %v", err)
		return err
	}

	logger.Log.Info("✅ Activity saved successfully")
	return nil
}
