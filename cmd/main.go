package main

import (
	"log"
	"project/config"
	"project/internal/bot"
)

func main() {
	// Загружаем переменные
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Error load .env: %v", err)
	}

	token := config.GetBotToken()
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN not installed")
	}

	bot.StartBot(token)
}
