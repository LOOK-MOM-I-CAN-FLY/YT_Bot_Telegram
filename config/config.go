package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("File .env not found, environment variables from the system are used")
	}
	return err
}

func GetBotToken() string {
	return os.Getenv("TELEGRAM_BOT_TOKEN")
}
