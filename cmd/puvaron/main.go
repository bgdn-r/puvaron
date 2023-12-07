package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bgdn-r/puvaron/pkg/config"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error: %v\n", err)
	}
}

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	r := chi.NewRouter()

	srv := &http.Server{
		Addr:    cfg.ListenAddr,
		Handler: r,
	}

	go func() {
		log.Printf("server listening on %s\n", srv.Addr)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("error: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v\n", err)
	}

	log.Println("server shutdown gracefully.")
}
