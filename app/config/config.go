package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rubikge/telegram_2_dropbox/internal/models"
)

func LoadConfig() *models.Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &models.Config{
		Telegram: models.TelegramConfig{
			BotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		},
		Dropbox: models.DropboxConfig{
			Path:         os.Getenv("DROPBOX_PATH"),
			AppKey:       os.Getenv("DROPBOX_APP_KEY"),
			AppSecret:    os.Getenv("DROPBOX_APP_SECRET"),
			RefreshToken: os.Getenv("DROPBOX_REFRESH_TOKEN"),
		},
	}
}
