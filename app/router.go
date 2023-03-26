package app

import (
	"imageupload/handler"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	router.MaxMultipartMemory = 5 << 20 //5MB
	router.POST("/upload", handler.UploadImageHandler)
	router.Run() // listen and serve on 0.0.0.0:8080
}
