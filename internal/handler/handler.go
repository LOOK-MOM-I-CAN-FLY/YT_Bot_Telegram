package handler

import (
	"log"
	"project/internal/downloader"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		msg := tgbotapi.NewMessage(message.Chat.ID, "Привет! Отправь мне ссылку на YouTube.")
		bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Неизвестная команда.")
		bot.Send(msg)
	}
}

func HandleYouTubeLink(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	url := message.Text
	chatID := message.Chat.ID

	bot.Send(tgbotapi.NewMessage(chatID, "Скачиваю видео..."))

	filePath, err := downloader.DownloadYouTubeVideo(url)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, "Ошибка при загрузке видео."))
		log.Println(err)
		return
	}

	video := tgbotapi.NewVideo(chatID, tgbotapi.FilePath(filePath))
	_, err = bot.Send(video)
	if err != nil {
		log.Println("Ошибка при отправке видео:", err)
	}
}
