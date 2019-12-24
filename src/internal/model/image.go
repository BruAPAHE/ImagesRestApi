package model

import "go.mongodb.org/mongo-driver/mongo"

type ImageModel struct {
	collection  *mongo.Collection
}

func NewImageModel() *ImageModel {
	return &ImageModel{}
}
