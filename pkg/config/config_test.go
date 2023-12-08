package config

import (
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	_, err := Read()
	if err == nil {
		t.Errorf("no env vars not set but read succeeded: %v", err)
	}

	os.Setenv("LISTEN_ADDR", ":9090")
	_, err = Read()
	if err == nil {
		t.Errorf("some env vars not set but read succeeded: %v", err)
	}

	os.Setenv("DB_URI", "example_uri")
	if err == nil {
		t.Errorf("some env vars not set but read succeeded: %v", err)
	}
	os.Setenv("JWT_SECRET", "example_jwt_secret")

	_, err = Read()
	if err != nil {
		t.Errorf("env vars are set but the read failed: %v", err)
	}
}
