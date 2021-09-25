package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	password string
	hash     string
}

func NewHashPassword(hash string) (Password, error) {
	return Password{hash: hash}, nil
}

func NewPassword(password string) (Password, error) {
	return Password{password: password}, nil
}

func (p Password) HashPassword() (string, error) {
	if p.hash != "" {
		return p.hash, nil
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(p.password), 14)
	if err != nil {
		return "", fmt.Errorf("Error hashing password")
	}
	return string(bytes), nil
}

func (p Password) ValidatePassowrd(password string) bool {
	hash := p.hash
	if hash == "" {
		h, err := p.HashPassword()
		if err != nil {
			return false
		}
		hash = h
	}
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
