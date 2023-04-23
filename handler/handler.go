package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"imageupload/dao"
	"imageupload/models"
	"imageupload/repository"
	"imageupload/service"
)

type ImageHandler struct {
	ImageService service.ImageService
}

func NewImageHandler(container *models.Container) *ImageHandler {
	// Initialize DAOs
	imageDao := dao.NewImageDao(container.MongoDB)

	s3Service := repository.NewS3Service("bucketforimageupload")

	// Initialize services
	imageService := service.NewImageService(imageDao, s3Service)

	// Initialize handlers
	imageHandler := &ImageHandler{
		ImageService: imageService,
	}

	return imageHandler
}

// swagger:route POST /upload PostUploadImage UploadImageHandler
//
// Uploads an image file to the server.
//
// This endpoint accepts a multipart/form-data request with a single "file" parameter
// that contains the image file to upload. The maximum allowed file size is 5 MB.
//
// Responses:
//   200: UploadImageResponse
//   400: BadRequestResponse
//   500: InternalServerErrorResponse

func (h *ImageHandler) PostUploadImage(c *gin.Context) {
	//getting the file
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "failed to recieve image"})
		return
	}

	if !service.IsImage(fileHeader) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "file is not an image"})
		return
	}

	if fileHeader.Size > 5<<20 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "file size should not be more than 5MB"})
		return
	}

	file, openErr := fileHeader.Open()
	if openErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "failed to open image"})
		return
	}

	h.ImageService.UploadImage(c, file, fileHeader)
}

func (h *ImageHandler) GetListImages(c *gin.Context) {
	//service.ListImages(c)
	h.ImageService.ListImages(c)
}

func (h *ImageHandler) DeleteImage(c *gin.Context) {
	imageName := c.Param("imageName")
	h.ImageService.DeleteImage(c, imageName)
}
