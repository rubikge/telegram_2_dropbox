package services

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type DropBoxService struct {
	token      string
	folderPath string
}

func NewDropBoxService(token string, folderPath string) *DropBoxService {
	return &DropBoxService{token: token, folderPath: folderPath}
}

func (dropbox *DropBoxService) UploadToDropbox(fileData *[]byte, filename string) error {
	fullPath := fmt.Sprintf("%s/%s.jpg", dropbox.folderPath, strings.ReplaceAll(filename, "/", ""))
	fmt.Println(fullPath)

	url := "https://content.dropboxapi.com/2/files/upload"
	req, _ := http.NewRequest("POST", url, bytes.NewReader(*fileData))

	req.Header.Set("Authorization", "Bearer "+dropbox.token)
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
