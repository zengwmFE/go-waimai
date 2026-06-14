package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

var allExts = map[string]bool{
	".jpg":  true,
	".png":  true,
	".jpeg": true,
	".webp": true,
}

const maxFileSize = 10 << 20 // 5MB
const uploadDir = "uploads"

func SaveFile(file *multipart.FileHeader) (string, error) {
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allExts[ext] {
		return "", fmt.Errorf("unsupported file type: %s", ext)
	}

	if file.Size > maxFileSize {
		return "", fmt.Errorf("this file is over size")
	}

	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", err
	}

	filename := uuid.New().String() + ext
	dst := filepath.Join(uploadDir, filename)
	src, err := file.Open()

	if err != nil {
		return "", err
	}

	defer src.Close()

	out, err := os.Create(dst)

	if err != nil {
		return "", err
	}
	defer out.Close()
	if _, err := io.Copy(out, src); err != nil {
		return "", err
	}

	return filename, nil
}
