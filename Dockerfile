# Base image
FROM golang:1.17-alpine

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Install any dependencies
RUN go get -d -v ./...
RUN go install -v ./...

# Expose port 8080
EXPOSE 8080

# Start the application
CMD ["go", "run", "main.go"]