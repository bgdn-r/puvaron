package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bgdn-r/puvaron/pkg/config"
	"github.com/bgdn-r/puvaron/pkg/router"
	"github.com/joho/godotenv"

	_ "github.com/bgdn-r/puvaron/pkg/logger"
	_ "github.com/lib/pq"
)

func init() {
	if err := godotenv.Load(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	conn, err := sql.Open("postgres", cfg.DBUri)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	_ = conn

	r := router.NewRouter()

	srv := &http.Server{
		Addr:    cfg.ListenAddr,
		Handler: r,
	}

	go func() {
		slog.Info(fmt.Sprintf("server listening on %s", srv.Addr))
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			slog.Error(err.Error())
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Info("server shutdown gracefully.")
}
