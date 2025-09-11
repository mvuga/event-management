package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashValue(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(passwordHash), err
}

func CheckPasswordHash(password, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
