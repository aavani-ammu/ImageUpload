package repository

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUploadImage(t *testing.T) {
	// Create a test file
	testFile, err := os.Open("C:/Users/aavan/imageupload/image/flower.jpg")
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer testFile.Close()

	// Create a mock gin.Context object
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Call the function with the test file
	s := s3Service{}
	s.UploadImage(c, testFile, &multipart.FileHeader{
		Filename: "flower.jpg",
	})

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	expectedBody := gin.H{"message": "Image successfully uploaded", "url": "https://bucketforimageupload.s3.amazonaws.com/test_image.jpg"}
	var responseBody gin.H
	err = json.Unmarshal(w.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	if !reflect.DeepEqual(responseBody, expectedBody) {
		t.Errorf("Expected response body %v but got %v", expectedBody, responseBody)
	}
}
