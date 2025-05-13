package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
	"github.com/mholt/archiver/v3"
)

func main() {
	sourceFolder := `C:\KK_Data\syncthing\obsidian main vault` // Folder you want to back up
	backupFolder := `C:\KK_Data\Meta\obsidian backups` // Folder where backups will be stored

	// Ensure the backup folder exists
	if err := os.MkdirAll(backupFolder, os.ModePerm); err != nil {
		fmt.Println("Error creating backup directory:", err)
		return
	}

	// Create the backup file name based on the current date
	dateString := time.Now().Format("2006-01-02_15-04-05")
	backupFile := filepath.Join(backupFolder, fmt.Sprintf("backup_%s.zip", dateString))

	// Create the backup using archiver
	err := archiver.Archive([]string{sourceFolder}, backupFile)
	if err != nil {
		fmt.Println("Error creating backup:", err)
		return
	}

	fmt.Printf("Backup created: %s\n", backupFile)

	// Manage backups: Keep at most 4 backups
	files, err := os.ReadDir(backupFolder)
	if err != nil {
		fmt.Println("Error reading backup directory:", err)
		return
	}

	var backupFiles []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) == ".zip" { // Change to .zip
			backupFiles = append(backupFiles, file.Name())
		}
	}

	sort.Slice(backupFiles, func(i, j int) bool {
		infoI, _ := os.Stat(filepath.Join(backupFolder, backupFiles[i]))
		infoJ, _ := os.Stat(filepath.Join(backupFolder, backupFiles[j]))
		return infoI.ModTime().Before(infoJ.ModTime())
	})

	if len(backupFiles) > 4 {
		oldestBackup := filepath.Join(backupFolder, backupFiles[0])
		if err := os.Remove(oldestBackup); err == nil {
			fmt.Printf("Deleted oldest backup: %s\n", oldestBackup)
		} else {
			fmt.Println("Error deleting oldest backup:", err)
		}
	}

	fmt.Println("Backup process complete.")
}
