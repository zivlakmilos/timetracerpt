package timetrace

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func findLatestFile(dirPath string) (string, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return "", err
	}

	var latestFolder string
	var latestModTime time.Time
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		if info.ModTime().After(latestModTime) {
			latestModTime = info.ModTime()
			latestFolder = entry.Name()
		}
	}

	if latestFolder == "" {
		return "", fmt.Errorf("error: '%s' is empty", dirPath)
	}

	return filepath.Join(dirPath, latestFolder), nil
}

func CheckStatus() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	recordsDir := filepath.Join(homeDir, ".timetrace", "records")

	latestRecordsDir, err := findLatestFile(recordsDir)
	if err != nil {
		return "", err
	}
	latestRecord, err := findLatestFile(latestRecordsDir)
	if err != nil {
		return "", err
	}

	record, err := ParseRecord(latestRecord)
	if err != nil {
		return "", err
	}

	if !record.End.IsZero() {
		return "off", nil
	}

	tag := ""
loop:
	for _, t := range record.Tags {
		switch t {
		case "coding":
			tag = " (c)"
			break loop
		case "meeting":
			tag = " (m)"
			break loop
		}
	}

	duration := time.Since(record.Start)
	status := fmt.Sprintf("%s%s - %02d:%02d", record.Project.Key, tag, int(duration.Hours()), int(duration.Minutes())%60)

	return status, nil
}
