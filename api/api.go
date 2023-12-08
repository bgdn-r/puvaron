package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/bgdn-r/puvaron/internal/db"
	"github.com/bgdn-r/puvaron/pkg/config"
	"github.com/go-chi/chi/v5"
)

type PuvaronAPI struct {
	queries *db.Queries
	config  *config.Config
	Router  *chi.Mux
}

func NewPuvaronAPI(conn *sql.DB, config *config.Config) *PuvaronAPI {
	r := chi.NewMux()
	r.Use(LoggingMiddleware)
	r.Use(RealIPMiddleware)
	r.Use(RecoverMiddleware)
	r.Use(JSONMiddleware)
	r.Use(TimeoutMiddleware(time.Minute))

	puvaronAPI := &PuvaronAPI{
		queries: db.New(conn),
		config:  config,
		Router:  r,
	}

	r.Post("/api/v1/user", puvaronAPI.CreateUser)
	r.Post("/api/v1/login", puvaronAPI.Login)
	r.Route("/api/v1/protected", func(r chi.Router) {
		r.Use(AuthMiddleware(puvaronAPI.config))
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		})
	})
	return puvaronAPI
}
