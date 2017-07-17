package main

import (
	_ "github.com/joho/godotenv/autoload"

	"time"
	"net/http"

	"github.com/Zhanat87/golang-gorilla/app"
)

func main() {
	hs, logger := app.SetupServer()

	go func() {
		logger.Printf("Listening on http://0.0.0.0%s\n", hs.Addr)

		if err := hs.ListenAndServe(); err != http.ErrServerClosed {
			logger.Fatal(err)
		}
	}()

	app.GracefulServer(hs, logger, 5 * time.Second)
}
