package downloader

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// DownloadYouTubeVideo загружает видео с YouTube по ссылке
func DownloadYouTubeVideo(url string) (string, error) {
	outputDir := "downloads"
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, os.ModePerm)
	}

	// Генерируем путь для сохранения видео
	outputTemplate := filepath.Join(outputDir, "%(title)s.%(ext)s")
	cmd := exec.Command("yt-dlp", "-o", outputTemplate, url)

	// Запускаем команду yt-dlp
	if output, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("error running yt-dlp: %v, output: %s", err, string(output))
	}

	// Находим загруженный файл
	files, err := filepath.Glob(filepath.Join(outputDir, "*"))
	if err != nil || len(files) == 0 {
		return "", fmt.Errorf("no files found in %s: %v", outputDir, err)
	}

	// Возвращаем путь к первому файлу
	return files[0], nil
}
