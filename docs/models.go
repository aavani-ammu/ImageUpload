package docs

import "mime/multipart"

// swagger:parameters UploadImageHandler
type UploadImageRequest struct {
	// The image file to upload
	// in: formData
	// required: true
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// swagger:response UploadImageResponse
type UploadImageResponse struct {
	// in: body
	Body struct {
		// The URL of the uploaded image.
		//
		// example: https://example.com/images/image.jpg
		URL string `json:"url"`
	}
}

// swagger:response BadRequestResponse
type BadRequestResponse struct {
	// in: body
	Body struct {
		// The error message.
		//
		// example: failed to receive image
		Message string `json:"message"`
	}
}

// swagger:response InternalServerErrorResponse
type InternalServerErrorResponse struct {
	// in: body
	Body struct {
		// The error message.
		//
		// example: failed to upload image
		Message string `json:"message"`
	}
}
