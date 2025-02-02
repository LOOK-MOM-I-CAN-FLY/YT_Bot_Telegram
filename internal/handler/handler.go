package handler

import (
	"fmt"
	"log"
	"net/url"
	"project/internal/downloader"
	"project/internal/storage"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// ValidateYouTubeURL checks  
func ValidateYouTubeURL(link string) bool {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return false
	}
	return parsedURL.Host == "www.youtube.com" || parsedURL.Host == "youtube.com" || parsedURL.Host == "youtu.be"
}

func HandleYouTubeLink(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	url := message.Text
	chatID := message.Chat.ID

	if !ValidateYouTubeURL(url) {
		errorMsg := "Invalid YouTube link. Please provide a valid URL."
		bot.Send(tgbotapi.NewMessage(chatID, errorMsg))
		log.Println(errorMsg)
		return
	}

	bot.Send(tgbotapi.NewMessage(chatID, "Downloading video..."))

	filePath, err := downloader.DownloadYouTubeVideo(url)
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to download video: %v", err)
		bot.Send(tgbotapi.NewMessage(chatID, errorMsg))
		log.Println(errorMsg)
		return
	}

	log.Printf("Video downloaded successfully: %s", filePath)

	video := tgbotapi.NewVideo(chatID, tgbotapi.FilePath(filePath))
	_, err = bot.Send(video)
	if err != nil {
		log.Printf("Failed to send video: %v", err)
		bot.Send(tgbotapi.NewMessage(chatID, "Error sending video."))
		return
	}

	bot.Send(tgbotapi.NewMessage(chatID, "Video sent successfully!"))

	if err := storage.CleanupDownloads(); err != nil {
		log.Printf("Failed to clean up downloads: %v", err)
	}
}
