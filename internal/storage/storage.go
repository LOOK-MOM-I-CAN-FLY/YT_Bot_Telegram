package storage

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func CleanupDownloads() error {
	outputDir := "downloads"
	files, err := filepath.Glob(filepath.Join(outputDir, "*"))
	if err != nil {
		return err
	}

	for _, file := range files {
		if err := os.Remove(file); err != nil {
			return fmt.Errorf("file deletion error %s: %v", file, err)
		}
	}
	return nil
}

func CheckYtDlp() error {
	_, err := exec.LookPath("yt-dlp")
	if err != nil {
		return fmt.Errorf("yt-dlp not found, install it before launching the bot")
	}
	return nil
}
