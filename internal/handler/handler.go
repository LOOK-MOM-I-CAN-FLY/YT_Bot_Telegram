package handler

import (
	"log"
	"project/internal/downloader"
	"project/internal/storage"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		msg := tgbotapi.NewMessage(message.Chat.ID, "Hi! Send me the YouTube link..")
		bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Unknown command.")
		bot.Send(msg)
	}
}

func HandleYouTubeLink(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	url := message.Text
	chatID := message.Chat.ID

	bot.Send(tgbotapi.NewMessage(chatID, "I'm downloading a video..."))

	filePath, err := downloader.DownloadYouTubeVideo(url)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, "Error when uploading video."))
		log.Println(err)
		return
	}

	video := tgbotapi.NewVideo(chatID, tgbotapi.FilePath(filePath))
	_, err = bot.Send(video)
	if err != nil {
		log.Println("Error when sending video:", err)
	}

	if err := storage.CleanupDownloads(); err != nil {
		log.Println("Error clearing downloads:", err)
	}
}
