package middleware

import (
	"errors"
	"net/http"

	"github.com/bgdn-r/puvaron/pkg/httputil"
)

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost || r.Method == http.MethodPut {
			if r.Header.Get("Content-Type") != "application/json" {
				httputil.WriteErr(w, errors.New("Content-Type must be 'application/json'"), http.StatusUnsupportedMediaType)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
