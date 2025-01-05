package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Файл .env не найден, используются переменные окружения из системы")
	}
	return err
}

func GetBotToken() string {
	return os.Getenv("TELEGRAM_BOT_TOKEN")
}
