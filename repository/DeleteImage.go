package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (s *s3Service) DeleteImage(imageName string) error {
	svc := s3.New(s.session)

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(imageName),
	})

	//if err != nil {
	//	return err
	//}
	//return nil

	if err != nil {
		// Check for specific errors related to permissions
		aerr, ok := err.(awserr.Error)
		if ok && aerr.Code() == "AccessDenied" {
			return fmt.Errorf("access denied. please check iam policy for delete permissions")
		}
		return err
	}

	return nil

}
