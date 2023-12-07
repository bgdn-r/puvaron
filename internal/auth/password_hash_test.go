package auth

import (
	"testing"
)

func TestHash(t *testing.T) {
	cfg := NewHashConfig()

	str := "RandomString123$!"
	hash, err := Hash(str, cfg)
	if err != nil {
		t.Errorf("failed to generate hash: %v\n", err)
	}
	if hash == "" {
		t.Errorf("returned hash is empty")
	}

	hash2, err := Hash(str, cfg)
	if err != nil {
		t.Errorf("failed to generate hash: %v\n", err)
	}
	if hash == hash2 {
		t.Errorf("first and second hash are the same but musn't be\n")
	}
}

func TestCompareWithHash(t *testing.T) {
	cfg := NewHashConfig()

	str := "RandomString123$!"
	hash, err := Hash(str, cfg)
	if err != nil {
		t.Errorf("failed to generate hash: %v\n", err)
	}

	match, err := CompareWithHash(str, hash, cfg)
	if err != nil {
		t.Errorf("failed to generate hash: %v\n", err)
	}
	if !match {
		t.Errorf("strings do not match but they should")
	}

	wrongPassword := "!$432gnirtSmodnaR"
	match, err = CompareWithHash(wrongPassword, hash, cfg)
	if err != nil {
		t.Errorf("failed to generate hash: %v\n", err)
	}
	if match {
		t.Errorf("strings do match but they shouldn't")
	}
}
