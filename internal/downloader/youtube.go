package downloader

import (
	"os"
	"os/exec"
	"path/filepath"
)

func DownloadYouTubeVideo(url string) (string, error) {
	outputDir := "downloads"
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, os.ModePerm)
	}

	outputPath := filepath.Join(outputDir, "%(title)s.%(ext)s")
	cmd := exec.Command("yt-dlp", "-o", outputPath, url)

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return outputPath, nil
}
