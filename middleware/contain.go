package middleware

import (
	"imageupload/models"

	"github.com/gin-gonic/gin"
)

func ContainerMiddleware(container *models.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("container", container)
		c.Next()
	}
}
