package db

import (
	"database/sql"
	"fmt"
	"log"

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
	s :=setting.LoadSettings()

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

	fmt.Println(dsn)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ Error opening DB: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("❌ Error pinging DB: %v", err)
	}

	log.Println("✅ Database connection established")
}
