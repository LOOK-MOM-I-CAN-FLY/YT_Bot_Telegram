package storage

import (
	"os"
	"path/filepath"
)

func CleanupDownloads() error {
	outputDir := "downloads"
	files, err := filepath.Glob(filepath.Join(outputDir, "*"))
	if err != nil {
		return err
	}

	for _, file := range files {
		os.Remove(file)
	}
	return nil
}
