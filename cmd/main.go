package main

import (
	"log"
	"project/config"
	"project/internal/bot"
)

func main() {
	token := config.GetBotToken()
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set")
	}
	bot.StartBot(token)
}
