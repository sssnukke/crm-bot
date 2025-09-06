package utils

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func SaveBase64Image(base64Image, uploadDir string) (string, error) {
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %v", err)
	}

	parts := strings.Split(base64Image, ",")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid base64 image")
	}

	imageData, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 image: %v", err)
	}

	var ext string
	switch {
	case strings.Contains(parts[0], "image/jpeg"):
		ext = ".jpg"
	case strings.Contains(parts[0], "image/png"):
		ext = ".png"
	case strings.Contains(parts[0], "image/gif"):
		ext = ".gif"
	default:
		return "", fmt.Errorf("unsupported image format")
	}

	filename := fmt.Sprintf("employee_%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(uploadDir, filename)

	if err := os.WriteFile(filePath, imageData, 0666); err != nil {
		return "", fmt.Errorf("failed to save image: %v", err)
	}

	return filename, nil
}
