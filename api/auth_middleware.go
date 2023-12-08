package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/bgdn-r/puvaron/pkg/config"
	"github.com/golang-jwt/jwt"
)

type ContextKey string

const (
	ContextUserIDKey ContextKey = "userID"
)

// AuthMiddleware creates a middleware that checks for a valid authentication token.
func AuthMiddleware(cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header required", http.StatusUnauthorized)
				return
			}

			headerParts := strings.Split(authHeader, " ")
			if len(headerParts) != 2 {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			tokenString := headerParts[1]
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(cfg.JWTSecret), nil
			})

			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), ContextUserIDKey, claims["user_id"])
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
			}
		})
	}
}
