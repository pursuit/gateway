package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pursuit/gateway/internal"
	"github.com/pursuit/gateway/internal/config"
)

func main() {
	defer log.Println("Shutdown the server success")

	handler := internal.NewServer(config.Instance("./internal/config"))

	server := &http.Server{
		Addr:         ":5003",
		Handler:      handler,
		ReadTimeout:  310 * time.Second,
		WriteTimeout: 310 * time.Second,
	}

	go func() {
		log.Println("listen to 5003")
		if err := server.ListenAndServe(); err != http.ErrServerClosed && err != nil {
			panic(err)
		}
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)

	log.Println("Server is ready")
	<-sigint

	log.Println("Shutting down the server")

	if err := server.Shutdown(context.Background()); err != nil {
		panic(err)
	}
}
