package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/rubikge/telegram_2_dropbox/internal/models"
)

type DropBoxService struct {
	config      *models.DropboxConfig
	accessToken string
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func getAccessToken(refreshToken, appKey, appSecret string) (string, error) {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	data.Set("client_id", appKey)
	data.Set("client_secret", appSecret)

	resp, err := http.PostForm("https://api.dropboxapi.com/oauth2/token", data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}

func (s *DropBoxService) updateAccessToken() {
	accessToken, err := getAccessToken(s.config.RefreshToken, s.config.AppKey, s.config.AppSecret)
	if err != nil {
		fmt.Println("Error refreshing access token:", err)
		return
	}

	s.accessToken = accessToken
	fmt.Println("Access token updated")
}

func NewDropBoxService(config *models.DropboxConfig) (*DropBoxService, error) {
	accessToken, err := getAccessToken(config.RefreshToken, config.AppKey, config.AppSecret)
	if err != nil {
		return nil, err
	}

	s := DropBoxService{
		config:      config,
		accessToken: accessToken,
	}

	go func() {
		for {
			time.Sleep(3 * time.Hour)
			s.updateAccessToken()
		}
	}()

	return &s, nil
}

func (d *DropBoxService) UploadToDropbox(fileData *[]byte, filename string) error {
	fullPath := fmt.Sprintf("%s/%s.jpg", d.config.Path, strings.ReplaceAll(filename, "/", ""))
	fmt.Println(fullPath)

	url := "https://content.dropboxapi.com/2/files/upload"
	req, _ := http.NewRequest("POST", url, bytes.NewReader(*fileData))

	req.Header.Set("Authorization", "Bearer "+d.accessToken)
	req.Header.Set("Dropbox-API-Arg", fmt.Sprintf(`{"path": "%s","mode":"add","autorename":true}`, fullPath))
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("dropbox upload failed: %s", body)
	}

	return nil
}
