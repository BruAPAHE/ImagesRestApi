package model



type ImagesContract interface {
	DownloadImages(url string) error
}