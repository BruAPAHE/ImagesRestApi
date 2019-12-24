package model

import (
	"ImagesRestApi/src/internal/database"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type ImageModel struct {
	DB *mongo.Collection
}

func NewImageModel() *ImageModel {

	return &ImageModel{
		DB: database.ImagesConnection.Collection("images"),
	}
}

func (i *ImageModel) DownloadImages(url string) error {
	// Get the images...
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	return nil
}
