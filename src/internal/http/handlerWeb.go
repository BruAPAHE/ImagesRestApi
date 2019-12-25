package http

import (
	"net/http"
)

type ViewData struct {
	Title   string
	Message string
}

func mainHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

//TODO: не дописал =(
	}
}
