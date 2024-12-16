package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func MakeRefreshToken() (string, error) {
	// Create a 32-byte array for the random data
	randomBytes := make([]byte, 32)

	// Read random data
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", errors.New("failed to generate random data:" + err.Error())
	}

	// Convert the random bytes to a hex-encoded string
	return hex.EncodeToString(randomBytes), nil
}
