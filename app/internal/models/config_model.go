package models

type TelegramConfig struct {
	BotToken string
}

type DropboxConfig struct {
	Path         string
	AppKey       string
	AppSecret    string
	RefreshToken string
}

type Config struct {
	Telegram TelegramConfig
	Dropbox  DropboxConfig
}
