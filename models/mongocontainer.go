package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Container defines the application container for managing dependencies
type Container struct {
	MongoClient *mongo.Client
	MongoDB     *mongo.Database
}

// NewContainer creates a new application container
func NewContainer() *Container {
	return &Container{}
}

// ConnectMongoDB connects to MongoDB and sets the database client in the container
func (c *Container) ConnectMongoDB(uri string) error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	db := client.Database("image_upload")

	c.MongoClient = client
	c.MongoDB = db

	return nil
}

// CloseMongoDB disconnects the MongoDB client
func (c *Container) CloseMongoDB() error {
	err := c.MongoClient.Disconnect(context.Background())
	if err != nil {
		return err
	}
	return nil
}
