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
	web.Router.HandleFunc("/v1/core/download/", MultipleMiddleware(imagesDownload(), middlewareList...)).Methods("GET")
	// WEB
}
