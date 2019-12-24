package main

import (
	app "ImagesRestApi/src/internal/application"
	"log"
)

func main() {
	Application := app.New()

	if err := Application.Start(); err != nil {
		log.Printf("Application didn't start:  %+v", err)
	}
}