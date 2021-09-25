package handler

import (
	"context"
	"encoding/json"

	auth "github.com/apascualco/apascualco-auth/internal"
	"github.com/google/uuid"
)

type AuthHandler struct {
	userRepository auth.UserRepository
}

func NewAuthHandler(ur auth.UserRepository) AuthHandler {
	return AuthHandler{userRepository: ur}
}

const QUEUE = "new.auth"

func (a AuthHandler) Queue() string {
	return QUEUE
}

type eventUser struct {
	Uuid        uuid.UUID `json:"uuid"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	AggregateID string    `json:"aggregateID"`
	OccurredOn  string    `json:"occurredOn"`
}

func mapToUser(eu eventUser) (auth.User, error) {
	uid := auth.NewUUID(eu.Uuid)
	email, err := auth.NewEmail(eu.Email)
	if err != nil {
		return auth.User{}, err
	}
	password, err := auth.NewHashPassword(eu.Password)
	if err != nil {
		return auth.User{}, err
	}
	return auth.NewUser(uid, email, password)
}

func (a AuthHandler) EventHandler(body string) error {

	var eu eventUser
	if err := json.Unmarshal([]byte(body), &eu); err != nil {
		return err
	}
	u, err := mapToUser(eu)
	if err != nil {
		return err
	}
	return a.userRepository.Save(context.Background(), u)
}
