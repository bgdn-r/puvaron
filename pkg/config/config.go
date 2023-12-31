package config

import (
	"fmt"
	"os"
)

type Config struct {
	ListenAddr string
	DBUri      string
	JWTSecret  string
}

func Read() (*Config, error) {
	listenAddr, ok := os.LookupEnv("LISTEN_ADDR")
	if !ok {
		return nil, fmt.Errorf("LISTEN_ADDR is required but was not set")
	}

	dbUri, ok := os.LookupEnv("DB_URI")
	if !ok {
		return nil, fmt.Errorf("DB_URI is required but was not set")
	}

	jwtSecret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		return nil, fmt.Errorf("JWT_SECRET is required but was not set")
	}

	return &Config{
		ListenAddr: listenAddr,
		DBUri:      dbUri,
		JWTSecret:  jwtSecret,
	}, nil
}
