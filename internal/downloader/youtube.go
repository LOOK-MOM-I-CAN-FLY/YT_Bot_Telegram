package downloader

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)


func DownloadYouTubeVideo(url string) (string, error) {
	outputDir := "downloads"
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(oututDir, os.ModePerm)
	}


	outputTemplate := filepath.Join(outputDir, "%(title)s.%(ext)s")
	cmd := exec.Command("yt-dlp", "-o", outputTemplate, url)


	if output, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("error running yt-dlp: %v, output: %s", err, string(output))
	}


	files, err := filepath.Glob(filepath.Join(outputDir, "*"))
	if err != nil || len(files) == 0 {
		return "", fmt.Errorf("no files found in %s: %v", outputDir, err)
	}


	return files[0], nil
}
