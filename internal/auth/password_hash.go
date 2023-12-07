package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type HashConfig struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

func NewHashConfig() *HashConfig {
	return &HashConfig{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}
}

func Hash(str string, cfg *HashConfig) (string, error) {
	salt := make([]byte, cfg.SaltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(str), salt, cfg.Iterations, cfg.Memory, cfg.Parallelism, cfg.KeyLength)

	return fmt.Sprintf("%s$%s", base64.RawStdEncoding.EncodeToString(hash), base64.RawStdEncoding.EncodeToString(salt)), nil
}

func CompareWithHash(str, hashedStr string, cfg *HashConfig) (bool, error) {
	parts := strings.Split(hashedStr, "$")
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid hashed str format")
	}
	salt, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, err
	}

	hash := argon2.IDKey([]byte(str), salt, cfg.Iterations, cfg.Memory, cfg.Parallelism, cfg.KeyLength)

	return parts[0] == base64.RawStdEncoding.EncodeToString(hash), nil
}
