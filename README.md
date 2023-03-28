# ImageUpload

A simple Go web application for uploading images to Amazon S3.

## Installation

To run this application, you need to have Go and Docker installed on your machine.

1. Clone this repository: git clone https://github.com/aavani-ammu/ImageUpload.git
2. Change into the directory: cd ImageUpload
3. Set your local environment with aws credentials in .env file and set your aws s3 bucket name for image upload in UploadImage function in repository/awsS3.go
3. Start the application: docker-compose up

## Usage

Once the application is running, you can access it at http://localhost:8080/upload. From there, you can select an image file to upload.

The uploaded image will be stored in the Amazon S3 bucket.
