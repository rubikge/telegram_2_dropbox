package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// incomingMessage := models.TelegramIncomingMessage{
	// 	UpdateID: 1234567890,
	// 	Message: struct {
	// 		MessageID int64 `json:"message_id"`
	// 		From      struct {
	// 			ID           int64  `json:"id"`
	// 			IsBot        bool   `json:"is_bot"`
	// 			FirstName    string `json:"first_name"`
	// 			LastName     string `json:"last_name"`
	// 			Username     string `json:"username"`
	// 			LanguageCode string `json:"language_code"`
	// 		} `json:"from"`
	// 		Chat struct {
	// 			ID        int64  `json:"id"`
	// 			FirstName string `json:"first_name"`
	// 			LastName  string `json:"last_name"`
	// 			Username  string `json:"username"`
	// 			Type      string `json:"type"`
	// 		} `json:"chat"`
	// 		Date  int64  `json:"date"`
	// 		Text  string `json:"text"`
	// 		Photo []struct {
	// 			FileID       string `json:"file_id"`
	// 			FileUniqueID string `json:"file_unique_id"`
	// 			FileSize     int64  `json:"file_size"`
	// 			Width        int64  `json:"width"`
	// 			Height       int64  `json:"height"`
	// 		} `json:"photo"`
	// 		Caption string `json:"caption"`
	// 	}{
	// 		MessageID: 9876543210,
	// 		From: struct {
	// 			ID           int64  `json:"id"`
	// 			IsBot        bool   `json:"is_bot"`
	// 			FirstName    string `json:"first_name"`
	// 			LastName     string `json:"last_name"`
	// 			Username     string `json:"username"`
	// 			LanguageCode string `json:"language_code"`
	// 		}{
	// 			ID:           1234567,
	// 			IsBot:        false,
	// 			FirstName:    "John",
	// 			LastName:     "Doe",
	// 			Username:     "johndoe",
	// 			LanguageCode: "en",
	// 		},
	// 		Chat: struct {
	// 			ID        int64  `json:"id"`
	// 			FirstName string `json:"first_name"`
	// 			LastName  string `json:"last_name"`
	// 			Username  string `json:"username"`
	// 			Type      string `json:"type"`
	// 		}{
	// 			ID:        7654321,
	// 			FirstName: "Jane",
	// 			LastName:  "Doe",
	// 			Username:  "janedoe",
	// 			Type:      "private",
	// 		},
	// 		Date: time.Now().Unix(),
	// 		Text: "Hello, this is a test message!",
	// 		Photo: []struct {
	// 			FileID       string `json:"file_id"`
	// 			FileUniqueID string `json:"file_unique_id"`
	// 			FileSize     int64  `json:"file_size"`
	// 			Width        int64  `json:"width"`
	// 			Height       int64  `json:"height"`
	// 		}{},
	// 		Caption: "Test photo caption",
	// 	},
	// }

	incomingMessage := struct{}{}
	jsonData, err := json.MarshalIndent(incomingMessage, "", "  ")
	if err != nil {
		return
	}

	buffer := bytes.NewBuffer(jsonData)

	resp, err := http.Post("https://967f-82-211-155-179.ngrok-free.app/webhook", "application/json", buffer)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Response status %d", resp.StatusCode)
}
