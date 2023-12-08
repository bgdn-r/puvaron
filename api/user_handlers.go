package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/bgdn-r/puvaron/internal/auth"
	"github.com/bgdn-r/puvaron/internal/db"
	"github.com/google/uuid"
)

func (a *PuvaronAPI) CreateUser(rw http.ResponseWriter, r *http.Request) {
	var cur CreateUserRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cur); err != nil {
		slog.Error(err.Error())
		WriteErr(rw, errors.New("failed to decode JSON"), http.StatusBadRequest)
		return
	}

	hashedPassword, err := auth.Hash(cur.Password, auth.NewHashConfig())
	if err != nil {
		slog.Error(err.Error())
		WriteErr(rw, errors.New("failed to hash password"), http.StatusBadRequest)
		return
	}

	_, err = a.queries.CreateUser(r.Context(), db.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FirstName: cur.FirstName,
		LastName:  cur.LastName,
		Username:  cur.Username,
		Email:     cur.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		slog.Error(err.Error())
		WriteErr(rw, errors.New("failed to create user"), http.StatusInternalServerError)
		return
	}

	resp := UserResponse{
		Email:     cur.Email,
		FirstName: cur.FirstName,
		LastName:  cur.LastName,
		Username:  cur.Username,
		Phone:     cur.Phone,
	}

	WriteJSON(rw, resp, http.StatusCreated)
}

func (a *PuvaronAPI) Login(rw http.ResponseWriter, r *http.Request) {
	var lr LoginRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&lr); err != nil {
		slog.Error(err.Error())
		WriteErr(rw, errors.New("failed to decode JSON"), http.StatusBadRequest)
		return
	}

	user, err := a.queries.GetUserByEmail(r.Context(), lr.Email)
	if err != nil {
		slog.Error(err.Error())
		WriteErr(rw, errors.New("not found"), http.StatusNotFound)
		return
	}

	ok, err := auth.CompareWithHash(lr.Password, user.Password, auth.NewHashConfig())
	if err != nil {
		slog.Error(err.Error())
		WriteErr(rw, err, http.StatusBadRequest)
		return
	}
	if !ok {
		WriteErr(rw, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	token, err := GenerateJWT(user.ID.String(), a.config)
	if err != nil {
		slog.Error(err.Error())
		WriteErr(rw, err, http.StatusBadRequest)
		return
	}

	resp := map[string]string{
		"access_token": token,
	}

	WriteJSON(rw, resp, http.StatusOK)
}
