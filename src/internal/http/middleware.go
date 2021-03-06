package http

import (
	"log"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

var middlewareList = []Middleware{
	existUrlParamsMiddleware,
	MainMiddleware,
}

// Main middleware....
func MultipleMiddleware(handle http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	// проверка на наличие middle..
	if len(middleware) < 1 {
		return handle
	}

	wrapped := handle

	// Проходим по каждому посреднику с конца
	for i := len(middleware) - 1; i >= 0; i-- {
		wrapped = middleware[i](wrapped)
	}

	return wrapped
}

func existUrlParamsMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Query().Get("url") == "" {
			response, err := MakeJSONResponse(http.StatusBadRequest, "Bad params: `url` is empty")
			if err != nil {
				log.Printf("Cannot make response: %+v", err)
			}

			writer.WriteHeader(http.StatusBadRequest)

			if _, err := writer.Write(response); err != nil {
				log.Printf("Cannot write response: %+v", err)
			}

			return
		}
		// next
		handler(writer, request)
	}
}

func MainMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		// next
		handler(writer, request)
	}
}
