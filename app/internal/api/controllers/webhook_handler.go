package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/rubikge/telegram_2_dropbox/internal/models"
	"github.com/rubikge/telegram_2_dropbox/internal/services"
)

type Webhook struct {
	telegramService *services.TelegramService
	dropboxService  *services.DropBoxService
	lastMessageText string
}

func NewWebhook(telegramService *services.TelegramService, dropboxService *services.DropBoxService) *Webhook {
	return &Webhook{telegramService: telegramService, dropboxService: dropboxService, lastMessageText: ""}
}

func (w *Webhook) Handler(c fiber.Ctx) error {
	var message models.TelegramIncomingMessage

	if err := c.Bind().JSON(&message); err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusOK)
	}

	if message.Message.Text != "" {
		w.lastMessageText = message.Message.Text
		return c.SendStatus(fiber.StatusOK)
	}

	photo, err := w.telegramService.GetPhoto(&message)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusOK)
	}

	fileName := fmt.Sprintf("без имени %d", message.Message.Date)

	if photo.Caption != "" {
		fileName = photo.Caption
	}

	if w.lastMessageText != "" {
		fileName = w.lastMessageText
	}

	err = w.dropboxService.UploadToDropbox(&photo.FileData, fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Photo successfully uploaded")
	}

	w.lastMessageText = ""
	return c.SendStatus(fiber.StatusOK)
}
