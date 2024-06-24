package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPass), nil
}

func ComparePassword(hashedPass []byte, password string) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	return err
}
