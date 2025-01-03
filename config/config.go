package config

import (
	"os"
)

func GetBotToken() string {
	return os.Getenv("TELEGRAM_BOT_TOKEN")
}
