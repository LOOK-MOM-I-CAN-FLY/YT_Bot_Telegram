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
	log.Printf("Logged in under an account %s", bot.Self.UserName)

	// Проверяем наличие yt-dlp
	if err := storage.CheckYtDlp(); err != nil {
		log.Fatalf("yt-dlp не найден: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("Received message: %v", update.Message.Text)

		// Check if the message is a command or a valid YouTube link
		if update.Message.IsCommand() || handler.ValidateYouTubeURL(update.Message.Text) {
			handler.HandleYouTubeLink(bot, update.Message)
		}
	}
}
