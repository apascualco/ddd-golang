package auth

import (
	"github.com/apascualco/apascualco-auth/kit/event"
	"github.com/google/uuid"
)

type User struct {
	uuid     UUID
	email    Email
	password Password

	events []event.Event
}

func NewUser(uuid UUID, email Email, password Password) (User, error) {
	return User{
		uuid:     uuid,
		email:    email,
		password: password,
	}, nil
}

func (u User) ID() uuid.UUID {
	return u.uuid.uuid
}

func (u User) Email() string {
	return u.email.String()
}

func (u User) ValidatePassword(password string) bool {
	return u.password.ValidatePassowrd(password)
}

func (u User) HashPassword() (string, error) {
	return u.password.HashPassword()
}

func (u User) P() Password {
	return u.password
}

func (u *User) Record(evt event.Event) {
	u.events = append(u.events, evt)
}

func (u User) PullEvents() []event.Event {
	evt := u.events
	u.events = []event.Event{}

	return evt
}
