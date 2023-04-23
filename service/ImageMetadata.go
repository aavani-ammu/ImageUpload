package service

import (
	"imageupload/models"
	"net/http"
	"time"
)

func createImageMetadata(url string, imageName string, fileSize int64) *models.ImageMetadata {
	return &models.ImageMetadata{
		ImageName:   imageName,
		ImageUrl:    url,
		ContentType: http.DetectContentType(nil),
		Size:        fileSize,
		CreatedAt:   time.Now(),
	}
}
