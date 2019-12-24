package repository

import "ImagesRestApi/src/internal/model"

type Images struct {
	Model model.ImagesContract
}

func NewImages(model model.ImagesContract) *Images {

	return &Images{
		Model: model,
	}
}

func (i *Images) DownloadImages(url string) error {
	
	return i.Model.DownloadImages(url)
}
