package service

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"

	rep "imageupload/repository"
)

func UploadImage(c *gin.Context, file multipart.File, fileHeader *multipart.FileHeader) {
	rep.UploadImage(c, file, fileHeader)

}
