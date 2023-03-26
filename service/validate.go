package service

import (
	"mime/multipart"
	"net/http"
	"strings"
)

func IsImage(file *multipart.FileHeader) bool {
	fileBytes, err := file.Open()
	if err != nil {
		return false
	}

	// Check the first 512 bytes to detect the file type
	buffer := make([]byte, 512)
	_, err = fileBytes.Read(buffer)
	if err != nil {
		return false
	}

	// Determine the content type of the file using the buffer
	fileType := http.DetectContentType(buffer)

	// Check if the content type is an image type
	if strings.HasPrefix(fileType, "image/") {
		return true
	}

	return false
}
