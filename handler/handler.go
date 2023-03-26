package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"imageupload/service"
)

// swagger:route POST /upload uploadImage UploadImageHandler
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

func UploadImageHandler(c *gin.Context) {
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

	service.UploadImage(c, file, fileHeader)
}
