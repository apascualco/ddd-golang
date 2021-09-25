package auth

import (
	"fmt"
	"net/mail"
)

type Email struct {
	email mail.Address
}

func NewEmail(email string) (Email, error) {
	a, err := mail.ParseAddress(email)
	if err != nil {
		return Email{}, fmt.Errorf("The email address is invalid")
	}
	return Email{email: *a}, nil
}

func (e Email) String() string {
	return e.email.Address
}
