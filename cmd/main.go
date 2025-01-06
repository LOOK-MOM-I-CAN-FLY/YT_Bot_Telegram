package main

import (
	"log"
	"project/config"
	"project/internal/bot"
)

func main() {
	// Загружаем переменные окружения
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Ошибка загрузки .env: %v", err)
	}

	token := config.GetBotToken()
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN не установлен")
	}

	bot.StartBot(token)
}
