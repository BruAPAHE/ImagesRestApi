package model

type MediaContract interface {
	DownloadByUrl(url string) error
}
