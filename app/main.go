package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/rubikge/telegram_2_dropbox/config"
	"github.com/rubikge/telegram_2_dropbox/internal/api"
	"github.com/rubikge/telegram_2_dropbox/internal/api/controllers"
	"github.com/rubikge/telegram_2_dropbox/internal/services"
)

func main() {
	config := config.LoadConfig()

	telegramService := services.NewTelegramService(config.TelegramBotToken)
	dropboxService := services.NewDropBoxService(config.DropboxAccessToken, config.DropboxPath)

	webhook := controllers.NewWebhook(telegramService, dropboxService)

	app := fiber.New()
	api.Router(app, webhook)
	app.Listen(":8080")
}
