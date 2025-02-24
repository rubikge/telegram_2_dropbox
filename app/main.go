package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/rubikge/telegram_2_dropbox/config"
	"github.com/rubikge/telegram_2_dropbox/internal/api"
	"github.com/rubikge/telegram_2_dropbox/internal/api/controllers"
	"github.com/rubikge/telegram_2_dropbox/internal/services"
)

func main() {
	config := config.LoadConfig()

	telegramService := services.NewTelegramService(&config.Telegram)
	dropboxService, err := services.NewDropBoxService(&config.Dropbox)
	if err != nil {
		log.Fatalln(err)
	}

	webhook := controllers.NewWebhook(telegramService, dropboxService)

	app := fiber.New()
	api.Router(app, webhook)
	app.Listen(":8080")
}
