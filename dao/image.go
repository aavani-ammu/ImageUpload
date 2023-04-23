package dao

import (
	"context"

	"imageupload/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ImageDao interface {
	CreateImage(image *models.ImageMetadata) error
	DeleteImage(imageName string) error
}

type imageDao struct {
	collection *mongo.Collection
}

func NewImageDao(db *mongo.Database) ImageDao {
	collection := db.Collection("images")
	return &imageDao{collection: collection}
}

func (dao *imageDao) CreateImage(image *models.ImageMetadata) error {
	_, err := dao.collection.InsertOne(context.Background(), image)
	if err != nil {
		return err
	}
	return nil
}

func (dao *imageDao) DeleteImage(imageName string) error {
	_, err := dao.collection.DeleteOne(context.Background(), bson.M{"name": imageName})
	if err != nil {
		return err
	}
	return nil
}
