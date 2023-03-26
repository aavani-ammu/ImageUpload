package rep

import (
	"context"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func SetupS3Uploader() *manager.Uploader {
	//set up s3 uploader
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	return uploader

}
func UploadImage(c *gin.Context, file multipart.File, fileHeader *multipart.FileHeader) {

	uploader := SetupS3Uploader()
	result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:  aws.String("bucketforimageupload"),
		Key:     aws.String(fileHeader.Filename),
		Body:    file,
		ACL:     "public-read",
		Expires: aws.Time(time.Now().Add(7 * 24 * time.Hour)),
	})

	if uploadErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "failed to upload image"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Image successfully uploaded", "url": result.Location})
}
