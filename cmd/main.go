package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pursuit/gateway/internal/proto/out/api/portal"
	"github.com/pursuit/gateway/internal/rest"

	"google.golang.org/grpc"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	defer log.Println("Shutdown the server success")

	portalConn, err := grpc.Dial("portal:5001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer portalConn.Close()

	userClient := portal_proto.NewUserClient(portalConn)

	userHandler := rest.Handler{
		UserClient: userClient,
	}
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))
	r.Mount("/metrics", promhttp.Handler())
	r.Post("/users", userHandler.CreateUser)
	r.Post("/user-sessions", userHandler.Login)

	server := &http.Server{
		Addr:    ":5003",
		Handler: r,
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
