package http

import (
	"ImagesRestApi/src/internal/model"
	"ImagesRestApi/src/internal/repository"
	"encoding/json"
	"log"
	"net/http"
)

func ImagesDownload() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var response ResponseBody
		var url string

		url = request.URL.Query().Get("url")

		imageModel := model.NewImageModel()

		repo := repository.NewImages(imageModel)

		if err := repo.DownloadByUrl(url); err != nil {
			response, err := MakeJSONResponse(http.StatusInternalServerError, err.Error())
			if err != nil {
				log.Printf("Cannot marshal to json: %v", err)
			}

			writer.WriteHeader(http.StatusNotFound)
			writer.Write(response)

			return
		}

		response.Meta.Code = http.StatusOK
		response.Meta.Message = "File was add to DB. Successful"

		if err := json.NewEncoder(writer).Encode(response); err != nil {
			writer.WriteHeader(500)
			return
		}
	}
}
