package scheduled

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func CleanupOldFiles(dir string, maxAge time.Duration) {
	log.Println("Cleaning up files in directory", dir)
	now := time.Now()

	count := 0

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatal("Failed to create the temp dir")
	}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing file %s: %v", path, err)
			return nil
		}
		if info.IsDir() {
			return nil
		}
		fileAge := now.Sub(info.ModTime())
		if fileAge > maxAge {
			if err := os.Remove(path); err != nil {
				log.Printf("Failed to delete file %s: %v", path, err)
			}
			count++
		}
		return nil
	})

	if err != nil {
		log.Printf("Error cleaning up files in directory %s: %v", dir, err)
	}
	log.Println("Deleted", count, "files")
}
