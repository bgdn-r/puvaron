package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	r := Response{
		StatusCode: statusCode,
		Data:       data,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(&r); err != nil {
		slog.Error(err.Error())
	}
}

func WriteErr(w http.ResponseWriter, err error, statusCode int) {
	r := Response{
		StatusCode: statusCode,
		Message:    err.Error(),
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(&r); err != nil {
		slog.Error(err.Error())
	}
}
