package app

import (
	router "ImagesRestApi/src/internal/http"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer *http.Server
	http       *router.Http
}

// Create Web instance ...
func New() *App {

	return &App{
		http: router.New(),
	}
}

// Start ...
func (app *App) Start() error {
	PORT := os.Getenv("APP_PORT")

	app.http.RegistrationRouters()

	app.httpServer = &http.Server{
		Addr:    PORT,
		Handler: app.http.Router,
	}

	go func() {
		log.Printf("Listening on %s...\n", PORT)
		if err := app.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return app.httpServer.Shutdown(ctx)
}
