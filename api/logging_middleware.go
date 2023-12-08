package api

import (
	"log/slog"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrappedWriter := NewWrappedResponseWriter(w)

		startTime := time.Now()

		next.ServeHTTP(wrappedWriter, r)

		duration := time.Since(startTime)
		slog.Info("http", slog.Group("request",
			slog.String("method", r.Method),
			slog.String("uri", r.RequestURI),
			slog.Int("status", wrappedWriter.StatusCode),
			slog.Duration("duration", duration),
		))
	})
}

// WrappedResponseWriter wraps an http.ResponseWriter to capture the status code written to it.
type WrappedResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

// NewWrappedResponseWriter creates a new WrappedResponseWriter.
func NewWrappedResponseWriter(w http.ResponseWriter) *WrappedResponseWriter {
	return &WrappedResponseWriter{ResponseWriter: w, StatusCode: http.StatusOK}
}

// WriteHeader captures the status code and calls the original WriteHeader method.
func (w *WrappedResponseWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
