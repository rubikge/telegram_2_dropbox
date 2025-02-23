package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/rubikge/telegram_2_dropbox/internal/api/controllers"
)

func Router(app *fiber.App, webhook *controllers.Webhook) {
	app.Post("/webhook", webhook.Handler)
}
