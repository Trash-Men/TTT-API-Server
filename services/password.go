package services

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct{}

func (_ PasswordService) Hash(password string) (string, error) {
	hash, error := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if error != nil {
		return "", error
	}
	return string(hash), nil
}

func (_ PasswordService) Validate(password string, hashedPassword string) error {
	if error := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); error != nil {
		return error
	}

	return nil
}
