package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//func ConnectMongo() {
// Connect to MongoDB
//client, err := ConnectMongoDB("mongodb://localhost:27017")
//if err != nil {
//	log.Fatal(err)
//}
//defer client.Disconnect(context.Background())
//}

// Database wraps a MongoDB database client
type Database struct {
	Client *mongo.Client
}

// Connect connects to MongoDB and returns a new Database instance
func Connect(uri string) (*Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &Database{Client: client}, nil
}

// Close disconnects the MongoDB client
func (d *Database) Close() error {
	err := d.Client.Disconnect(context.Background())
	if err != nil {
		return err
	}
	return nil
}
