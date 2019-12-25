package http

import (
	"github.com/gorilla/mux"
)

type Http struct {
	Router *mux.Router
}

func New() *Http {

	return &Http{
		Router: mux.NewRouter(),
	}
}

// Router List ...
func (web *Http) RegistrationRouters() {
	// API
	web.Router.HandleFunc("/api/v1/media/download/", MultipleMiddleware(ImagesDownload(), middlewareList...)).Methods("POST")
	// WEB
	web.Router.HandleFunc("/main", mainHandler())
}
