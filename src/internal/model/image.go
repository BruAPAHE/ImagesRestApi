package model

import (
	"ImagesRestApi/src/internal/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type ImageModel struct {
	DB *mongo.Database
}

func NewImageModel() *ImageModel {

	return &ImageModel{
		DB: database.ImagesConnection,
	}
}

func (i *ImageModel) DownloadByUrl(url string) error {
	var body []byte
	// Get the images...
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Printf("Can't close response body: %v", err)
		}
	}()

	if response.StatusCode == http.StatusOK {
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
	}

	bucket, err := gridfs.NewBucket(i.DB)
	if err != nil {
		return err
	}

	uploadStream, err := bucket.OpenUploadStream("images_" + strconv.Itoa(rand.Int()))
	if err != nil {
		return err
	}

	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(body)
	if err != nil {
		return err
	}

	log.Printf("Write file to DB was successful. File size: %d \n", fileSize)

	return nil
}
