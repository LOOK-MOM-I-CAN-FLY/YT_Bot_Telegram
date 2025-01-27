package downloader

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)


func DownloadYouTubeVideo(url string) (string, error) {
	outputDir := "downloads"
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, os.ModePerm)
	}


	outputTemplate := filepath.Join(outputDir, "%(title)s.%(ext)s")
	cmd := exec.Command("yt-dlp", "-o", outputTemplate, url)


	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cmd = exec.CommandContext(ctx, "yt-dlp", "-o", outputTemplate, url)


	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running yt-dlp: %v, output: %s", err, string(output))
	}


	files, err := filepath.Glob(filepath.Join(outputDir, "*"))
	if err != nil || len(files) == 0 {
		return "", fmt.Errorf("no files found in %s: %v", outputDir, err)
	}


	return files[0], nil
}
