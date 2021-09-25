package signin

import (
	"context"
	"fmt"

	auth "github.com/apascualco/apascualco-auth/internal"
	"github.com/apascualco/apascualco-auth/kit/event"
	"github.com/google/uuid"
)

type Signin struct {
	repository auth.UserRepository
	eventBus   event.Bus
}

func NewSignin(repository auth.UserRepository, eventBus event.Bus) Signin {
	return Signin{repository: repository, eventBus: eventBus}
}

func (s Signin) Signin(ctx context.Context, email, password string) error {
	e, err := auth.NewEmail(email)
	if err != nil {
		return err
	}
	if password == "" {
		return fmt.Errorf("The password should be null or empty for user %s", email)
	}
	p, err := auth.NewPassword(password)
	if err != nil {
		return err
	}
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	uid := auth.NewUUID(id)
	user, err := auth.NewUser(uid, e, p)
	if err != nil {
		return err
	}
	hp, _ := user.HashPassword()
	user.Record(NewSigninEvent(user.ID().String(), user.Email(), hp))
	err = s.repository.Save(ctx, user)
	if err != nil {
		return err
	}
	s.eventBus.Publish(ctx, user.PullEvents())
	return nil
}
