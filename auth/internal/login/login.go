package login

import (
	"context"
	"time"

	auth "github.com/apascualco/apascualco-auth/internal"
	"github.com/apascualco/apascualco-auth/kit/event"
	"github.com/dgrijalva/jwt-go"
)

type Login struct {
	repository auth.UserRepository
	eventBus   event.Bus
	secret     string
}

func NewLogin(repository auth.UserRepository, eventBus event.Bus, secret string) Login {
	return Login{repository: repository, eventBus: eventBus, secret: secret}
}

func (l Login) Login(ctx context.Context, email, password string) (string, error) {
	user, err := l.repository.SearchUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	ok := user.ValidatePassword(password)

	if ok {
		t, err := l.token(user.ID().String())
		if err != nil {
			return "", err
		}
		if t != "" {
			user.Record(NewLoginEvent(user.ID().String(), user.Email()))
			l.eventBus.Publish(ctx, user.PullEvents())
		}
		return t, nil
	}
	return "", nil
}

func (l Login) token(uuid string) (string, error) {
	secret := l.secret
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = uuid
	atClaims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
