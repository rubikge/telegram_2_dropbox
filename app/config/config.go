package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramBotToken   string
	DropboxAccessToken string
	DropboxPath        string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		TelegramBotToken:   os.Getenv("TELEGRAM_BOT_TOKEN"),
		DropboxAccessToken: os.Getenv("DROPBOX_ACCESS_TOKEN"),
		DropboxPath:        os.Getenv("DROPBOX_PATH"),
	}
}
