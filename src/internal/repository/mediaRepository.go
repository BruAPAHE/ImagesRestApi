package repository

import "ImagesRestApi/src/internal/model"

type Media struct {
	Model model.MediaContract
}

func NewImages(model model.MediaContract) *Media {

	return &Media{
		Model: model,
	}
}

func (i *Media) DownloadByUrl(url string) error {

	return i.Model.DownloadByUrl(url)
}
