package api

import (
	"time"

	"github.com/bgdn-r/puvaron/pkg/config"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(userID string, cfg *config.Config) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte(cfg.JWTSecret))
}
