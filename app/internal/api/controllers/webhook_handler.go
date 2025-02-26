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
	lastMessages    map[int64]string
}

func NewWebhook(telegramService *services.TelegramService, dropboxService *services.DropBoxService) *Webhook {
	return &Webhook{
		telegramService: telegramService,
		dropboxService:  dropboxService,
		lastMessages:    map[int64]string{},
	}
}

func (w *Webhook) Handler(c fiber.Ctx) error {
	var message models.TelegramIncomingMessage

	if err := c.Bind().JSON(&message); err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusBadGateway)
	}

	if txt := message.Message.Text; txt != "" {
		w.lastMessages[message.Message.Date] = txt
		return c.SendStatus(fiber.StatusOK)
	}

	photo, err := w.telegramService.GetPhoto(&message)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusBadGateway)
	}

	fileName := fmt.Sprintf("без имени %d", message.Message.Date)

	if txt := message.Message.Caption; txt != "" {
		fileName = txt
	}

	if txt := w.lastMessages[message.Message.Date]; txt != "" {
		fileName = txt
		delete(w.lastMessages, message.Message.Date)
	}

	err = w.dropboxService.UploadToDropbox(&photo.FileData, fileName)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusBadGateway)
	}

	fmt.Println("Photo successfully uploaded")
	return c.SendStatus(fiber.StatusOK)
}
