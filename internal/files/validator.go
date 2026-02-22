package files

import (
	"errors"
	"path/filepath"
)

// ValidateExtension ensures the file has a supported extension
func ValidateExtension(path string, allowedExts []string) error {
	ext := filepath.Ext(path)
	for _, allowed := range allowedExts {
		if ext == allowed {
			return nil
		}
	}
	return errors.New("unsupported file type: " + ext)
}

// ValidateExists ensures the file exists
func ValidateExists(path string) error {
	if !Exists(path) {
		return errors.New("file does not exist: " + path)
	}
	return nil
}