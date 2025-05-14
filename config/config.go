package config

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type Setting struct {
	Port     string
	Env      string
    DBPort   string
	DBName   string
	DBHost   string
	DBUser   string
	DBPass   string
	DBSSL    string
	BotToken string
	ChatID   string
}

func LoadSettings() Setting {
	
	viper.SetConfigFile(".env")
    viper.SetConfigType("env") // Ensure config type is set to 'env' to read the .env file

    if err := viper.ReadInConfig(); err != nil {
        log.Println("⚠️ .env file not found, falling back to system environment variables")
    } else {
        log.Println("✅ .env file loaded successfully")
    }

	sslmode := viper.GetString("DB_SSL")
	if sslmode == "" {
		sslmode = "disable"
	}

	return Setting{
		Port:     viper.GetString("PORT"),
		Env:      viper.GetString("ENV"),
		DBPort:   viper.GetString("DB_PORT"),
		DBHost:   viper.GetString("DB_HOST"),
		DBUser:   viper.GetString("DB_USER"),
		DBPass:   viper.GetString("DB_PASS"),
		DBSSL:    sslmode, // Correctly assigning sslmode
		BotToken: viper.GetString("TELEGRAM_BOT_TOKEN"),
		ChatID:   viper.GetString("TELEGRAM_CHAT_ID"),
	}
}
