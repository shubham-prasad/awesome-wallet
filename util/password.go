package util

import (
	"golang.org/x/crypto/bcrypt"
)

func CreatePasswordHash(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(hash string, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
}
