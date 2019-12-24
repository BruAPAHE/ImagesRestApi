package http

import (
	"ImagesRestApi/src/internal/model"
	"ImagesRestApi/src/internal/repository"
	"log"
	"net/http"
)

func imagesDownload() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var response []byte
		var url string


		imageModel := model.NewImageModel()

		repo := repository.NewImages(imageModel)

		if err := repo.DownloadImages(); err != nil {
			response, err := MakeJSONResponse(http.StatusInternalServerError, err.Error())
			if err != nil {
				log.Printf("Cannot marshal to json: %v", err)
			}

			writer.WriteHeader(http.StatusNotFound)
			writer.Write(response)
		}

		response, err := MakeJSONResponse(http.StatusOK, "OK")
		if err != nil {
			log.Printf("Cannot marshal to json: %v", err)
		}

		writer.WriteHeader(http.StatusOK)
		writer.Write(response)
	}
}
