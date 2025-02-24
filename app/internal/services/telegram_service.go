package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rubikge/telegram_2_dropbox/internal/models"
)

type TelegramService struct {
	config *models.TelegramConfig
}

func NewTelegramService(config *models.TelegramConfig) *TelegramService {
	return &TelegramService{config: config}
}

func (ts *TelegramService) GetPhoto(message *models.TelegramIncomingMessage) (*models.TelegramPhoto, error) {
	if len(message.Message.Photo) == 0 {
		return nil, fmt.Errorf("no photo in message")
	}

	fileID := message.Message.Photo[len(message.Message.Photo)-1].FileID
	fileData, err := downloadPhoto(fileID, ts.config.BotToken)
	if err != nil {
		return nil, err
	}

	return &models.TelegramPhoto{
		FileData: fileData,
		Caption:  message.Message.Caption,
	}, nil
}

func downloadPhoto(fileID string, token string) ([]byte, error) {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/getFile?file_id=%s", token, fileID)
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result models.TelegramFileResponse

	json.NewDecoder(resp.Body).Decode(&result)
	if !result.OK {
		return nil, fmt.Errorf("failed to get file path from Telegram")
	}

	fileURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", token, result.Result.FilePath)
	fileResp, err := http.Get(fileURL)
	if err != nil {
		return nil, err
	}
	defer fileResp.Body.Close()

	return io.ReadAll(fileResp.Body)
}
