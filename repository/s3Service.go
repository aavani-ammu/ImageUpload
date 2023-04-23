package repository

import (
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

type S3Service interface {
	UploadImage(c *gin.Context, file multipart.File, fileHeader *multipart.FileHeader) (string, error)
	ListImages(c *gin.Context)
	DeleteImage(imageName string) error
}

type s3Service struct {
	s3Client   *s3.S3
	bucketName string
	session    *session.Session
}

// NewS3Service creates a new S3Service instance with the given bucket name
func NewS3Service(bucketName string) *s3Service {

	// Create a new AWS session
	region := os.Getenv("AWS_REGION")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		panic(err)
	}

	// Create a new S3 client with the session
	s3Client := s3.New(sess)

	return &s3Service{
		s3Client:   s3Client,
		bucketName: bucketName,
		session:    sess,
	}
}
