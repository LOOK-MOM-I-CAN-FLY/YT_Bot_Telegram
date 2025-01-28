package downloader

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

//GENERAL FUNCTION
func DownloadYouTubeVideo(url string) (string, error) {
	outputDir := "downloads"
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, os.ModePerm)
	}

	outputTemplate := filepath.Join(outputDir, "%(title)s.%(ext)s")
	maxRetries := 3
	var err error
	var output []byte

	for i := 0; i < maxRetries; i++ {
		cmd := exec.Command("yt-dlp", "-o", outputTemplate, "--proxy", "http://your-proxy-address:port", url)
		output, err = cmd.CombinedOutput()
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second) // Wait before retrying
	}

	if err != nil {
		return "", fmt.Errorf("error running yt-dlp: %v, output: %s", err, string(output))
	}

	files, err := filepath.Glob(filepath.Join(outputDir, "*"))
	if err != nil || len(files) == 0 {
		return "", fmt.Errorf("no files found in %s: %v", outputDir, err)
	}

	return files[0], nil
}
