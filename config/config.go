package config

import (
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

	API_PREFIX_V1 string
	API_PREFIX_V2 string

	CORS_ALLOWED_ORIGINS string
}

func LoadSettings() Setting {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.SetDefault("API_PREFIX_V1", "/api/v1")
	viper.SetDefault("API_PREFIX_V2", "/api/v2")
	viper.SetDefault("CORS_ALLOWED_ORIGINS", "*")

	// Load .env if exisclear
	_ = viper.ReadInConfig()
	viper.AutomaticEnv()

	return Setting{
		Port	             : viper.GetString("PORT"),
		Env		             : viper.GetString("ENV"),
		DBPort	             : viper.GetString("DB_PORT"),
		DBHost	             : viper.GetString("DB_HOST"),
		DBUser	             : viper.GetString("DB_USER"),
		DBPass	             : viper.GetString("DB_PASS"),
		DBName	             : viper.GetString("DB_NAME"),
		DBSSL	             : viper.GetString("DB_SSL"),
		BotToken             : viper.GetString("TELEGRAM_BOT_TOKEN"),
		ChatID	             : viper.GetString("TELEGRAM_CHAT_ID"),
		API_PREFIX_V1        : viper.GetString("API_PREFIX_V1"),
		API_PREFIX_V2        : viper.GetString("API_PREFIX_V2"),
		CORS_ALLOWED_ORIGINS : viper.GetString("CORS_ALLOWED_ORIGINS"),
	}
}
