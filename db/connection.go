package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"entgo.io/ent/dialect"

	setting "logger/config"

	_ "github.com/lib/pq"
)

var (
	DB        *sql.DB
)

type DBSettings struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func InitDB() {
	s := setting.LoadSettings()

	db := DBSettings{
		Host		: s.DBHost,
		Port		: s.DBPort,
		User		: s.DBUser,
		Password	: s.DBPass,
		Name		: s.DBName,
		SSLMode		: s.DBSSL,
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		db.Host, db.Port, db.User, db.Password, db.Name, db.SSLMode,
	)

	var err error

	DB, err = sql.Open(dialect.Postgres, dsn)
	if err != nil {
		log.Fatalf("❌ Error opening DB: %v", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)
	DB.SetConnMaxIdleTime(1 * time.Minute)

	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		err = DB.Ping()
		if err == nil {
			log.Printf("✅ Database connection established (SSL Mode: %s)", db.SSLMode)
			return
		}
		log.Printf("⚠️ Attempt %d/%d: Failed to connect to database: %v", i+1, maxRetries, err)
		time.Sleep(time.Second * time.Duration(i+1))
	}

	log.Fatalf("❌ Failed to connect to database after %d attempts", maxRetries)
}
