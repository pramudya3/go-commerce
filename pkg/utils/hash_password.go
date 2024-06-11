package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass []byte) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		log.Printf("Hashing password failed, err: %v", err)
		return "", err
	}
	return string(hashed), nil
}
