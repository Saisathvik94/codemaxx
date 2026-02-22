package files

import (
	"fmt"
	"io"
	"os"
)

// BackupFile creates a backup with timestamp
func BackupFile(path string) (string, error) {
	if !Exists(path) {
		return "", fmt.Errorf("file does not exist to backup: %s", path)
	}

	backupPath := fmt.Sprintf("%s.bak", path)
	srcFile, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(backupPath)
	if err != nil {
		return "", err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return "", err
	}

	return backupPath, nil
}
