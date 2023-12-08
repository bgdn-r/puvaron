package api

import (
	"net/http"
	"strings"
)

// RealIPMiddleware sets the real IP address of the client to the request context.
func RealIPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ip := realIP(r); ip != "" {
			r.RemoteAddr = ip
		}
		next.ServeHTTP(w, r)
	})
}

// realIP extracts the real IP address from request headers.
func realIP(r *http.Request) string {
	if xForwardedFor := r.Header.Get("X-Forwarded-For"); xForwardedFor != "" {
		i := strings.Index(xForwardedFor, ", ")
		if i == -1 {
			return xForwardedFor
		}
		return xForwardedFor[:i]
	}

	return r.Header.Get("X-Real-IP")
}
