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
}

func LoadSettings() Setting {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	return Setting{
		Port:     viper.GetString("PORT"),
		Env:      viper.GetString("ENV"),
		DBPort:   viper.GetString("DB_PORT"),
		DBHost:   viper.GetString("DB_HOST"),
		DBUser:   viper.GetString("DB_USER"),
		DBPass:   viper.GetString("DB_PASS"),
		DBName:   viper.GetString("DB_NAME"),
		DBSSL:    viper.GetString("DB_SSL"),
		BotToken: viper.GetString("TELEGRAM_BOT_TOKEN"),
		ChatID:   viper.GetString("TELEGRAM_CHAT_ID"),
	}
}
