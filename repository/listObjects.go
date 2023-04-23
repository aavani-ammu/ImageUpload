package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"net/http"
)


// ListImages is a handler function that retrieves a list of images from an S3 bucket
func (s *s3Service) ListImages(c *gin.Context) {
	// Create a new S3 client with the session
	svc := s3.New(s.session)

	// Create a list objects input with the bucket name
	input := &s3.ListObjectsInput{
		Bucket: aws.String(s.bucketName),
	}

	// Get the list of objects in the bucket
	result, err := svc.ListObjects(input)
	if err != nil {
		// If there was an error retrieving the list, return the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a slice of strings to store the image keys
	imageKeys := []string{}

	// Loop through the objects and add the key to the image keys slice
	for _, item := range result.Contents {
		imageKeys = append(imageKeys, *item.Key)
	}

	// Return the list of image keys as JSON
	c.JSON(http.StatusOK, gin.H{"images": imageKeys})
}
