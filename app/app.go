package app

import (
	"os"
	"time"
	"os/signal"
	"syscall"
	"context"
	"log"
	"net/http"

	"github.com/Zhanat87/golang-gorilla/server"
)

func SetupServer() (*http.Server, *log.Logger) {
	logger := log.New(os.Stdout, "", 0)

	return &http.Server{
		Addr:    Port(),
		Handler: server.New(server.Logger(logger)),
	}, logger
}

func GracefulServer(hs *http.Server, logger *log.Logger, timeout time.Duration) {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	logger.Printf("\nShutdown with timeout: %s\n", timeout)

	if err := hs.Shutdown(ctx); err != nil {
		logger.Printf("Error: %v\n", err)
	} else {
		logger.Println("Server stopped")
	}
}

func Port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
