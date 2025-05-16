package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	setting "logger/config"

	_ "github.com/lib/pq"
)

// Global database connection
var DB *sql.DB

// TelegramSettings holds the Telegram bot configuration.
type TelegramSettings struct {
	BotToken string
	ChatID   string
}

// DBSettings holds the database configuration.
type DBSettings struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// Global config variables
var (
	db DBSettings
)

// InitDB initializes the DB connection from env or .env file
func InitDB() {
	// call config
	s := setting.LoadSettings()

	// Read DB configuration
	db = DBSettings{
		Host:     s.DBHost,
		Port:     s.DBPort,
		User:     s.DBUser,
		Password: s.DBPass,
		Name:     s.DBName,
		SSLMode:  s.DBSSL,
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		db.Host, db.Port, db.User, db.Password, db.Name, db.SSLMode,
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ Error opening DB: %v", err)
	}

	// Configure connection pool
	DB.SetMaxOpenConns(25)                 // Maximum number of open connections
	DB.SetMaxIdleConns(5)                  // Maximum number of idle connections
	DB.SetConnMaxLifetime(5 * time.Minute) // Maximum lifetime of a connection
	DB.SetConnMaxIdleTime(1 * time.Minute) // Maximum idle time of a connection

	// Try to connect with retries
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		err = DB.Ping()
		if err == nil {
			log.Printf("✅ Database connection established (SSL Mode: %s)", db.SSLMode)
			return
		}
		log.Printf("⚠️ Attempt %d/%d: Failed to connect to database: %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			time.Sleep(time.Second * time.Duration(i+1)) // Exponential backoff
		}
	}

	log.Fatalf("❌ Failed to connect to database after %d attempts", maxRetries)
}
