package bot

import (
	"log"
	"project/internal/handler"
	"project/internal/storage"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Logged in with an account %s", bot.Self.UserName)


	if err := storage.CheckYtDlp(); err != nil {
		log.Fatalf("yt-dlp not found: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("Received message: %v", update.Message.Text)

		if update.Message.IsCommand() {
			handler.HandleCommand(bot, update.Message)
		} else if update.Message.Text != "" {
			handler.HandleYouTubeLink(bot, update.Message)
		}
	}
}
