package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/bgdn-r/puvaron/pkg/config"
	"github.com/bgdn-r/puvaron/pkg/httputil"
	"github.com/golang-jwt/jwt"
)

// AuthMiddleware creates a middleware that checks for a valid authentication token.
func AuthMiddleware(cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authenticated, err := Authenticate(r, cfg)
			if err != nil {
				httputil.WriteErr(w, errors.New("internal server error"), http.StatusInternalServerError)
				return
			}
			if !authenticated {
				httputil.WriteErr(w, errors.New("unauthorized"), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// Authenticate verifies the JWT token from the Authorization header.
func Authenticate(r *http.Request, cfg *config.Config) (bool, error) {
	tokenString := extractToken(r)
	if tokenString == "" {
		return false, nil // No token provided
	}

	token, err := verifyToken(tokenString, cfg)
	if err != nil {
		return false, err // Token is invalid
	}

	return token.Valid, nil
}

// extractToken extracts the JWT token from the Authorization header.
func extractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// verifyToken validates the JWT token.
func verifyToken(tokenString string, cfg *config.Config) (*jwt.Token, error) {
	secretKey := []byte(cfg.JWTSecret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
