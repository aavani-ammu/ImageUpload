package main

import (
	"imageupload/handler"
	"imageupload/middleware"

	_ "imageupload/docs"
	"imageupload/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	container := models.NewContainer()
	defer container.CloseMongoDB()

	if err := container.ConnectMongoDB("mongodb://localhost:27017"); err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	// Initialize handlers
	imageHandler := handler.NewImageHandler(container)

	router := gin.Default()
	router.MaxMultipartMemory = 5 << 20 //5MB
	router.Use(middleware.ContainerMiddleware(container))

	router.POST("/upload", imageHandler.PostUploadImage)
	router.GET("/images", imageHandler.GetListImages)
	router.DELETE("/delete/:imageName", imageHandler.DeleteImage)

	// listen and serve on 0.0.0.0:8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
