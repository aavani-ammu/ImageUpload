package service

import (
	"imageupload/dao"
	"imageupload/models"
	"imageupload/repository"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ImageService defines the interface for image-related operations
type ImageService interface {
	UploadImage(c *gin.Context, file multipart.File, fileHeader *multipart.FileHeader) (*models.ImageMetadata, error)
	ListImages(c *gin.Context)
	DeleteImage(c *gin.Context, imageName string)
}

// imageService is a struct that implements ImageService interface
type imageService struct {
	imageDao  dao.ImageDao
	s3Service repository.S3Service
}

// NewImageService creates a new image service
func NewImageService(imageDao dao.ImageDao, s3Service repository.S3Service) ImageService {
	return &imageService{
		imageDao:  imageDao,
		s3Service: s3Service,
	}
}

// UploadImage uploads an image and stores its metadata
func (imageservice *imageService) UploadImage(c *gin.Context, file multipart.File, fileHeader *multipart.FileHeader) (*models.ImageMetadata, error) {
	// Upload image to S3

	imageUrl, err := imageservice.s3Service.UploadImage(c, file, fileHeader)
	if err != nil {
		return nil, err
	}

	// Create an instance of ImageMetaData struct using createImageMetadata function
	metadata := createImageMetadata(imageUrl, fileHeader.Filename, fileHeader.Size)

	err = imageservice.imageDao.CreateImage(metadata)
	if err != nil {
		return nil, err
	}

	return metadata, nil
}

func (imageservice *imageService) ListImages(c *gin.Context) {
	imageservice.s3Service.ListImages(c)
}

func (imageservice *imageService) DeleteImage(c *gin.Context, imageName string) {

	// Delete image from S3 bucket
	if err := imageservice.s3Service.DeleteImage(imageName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "failed to delete image from S3"})
		return
	}

	// Delete image metadata from MongoDB
	if err := imageservice.imageDao.DeleteImage(imageName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "failed to delete image metadata from database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image successfully deleted"})
}
