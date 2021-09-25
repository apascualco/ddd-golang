package signin

import (
	"context"
	"errors"

	"github.com/apascualco/apascualco-auth/kit/command"
)

const SiginCommandType command.Type = "apascualco.1.query.sigin"

type SigninCommand struct {
	email    string
	password string
}

func NewSigninCommand(email, password string) SigninCommand {
	return SigninCommand{email: email, password: password}
}

func (q SigninCommand) Type() command.Type {
	return SiginCommandType
}

type SigninCommandHandler struct {
	signin Signin
}

func NewSigninCommandHandler(s Signin) SigninCommandHandler {
	return SigninCommandHandler{signin: s}
}

func (h SigninCommandHandler) Handle(ctx context.Context, c command.Command) error {
	sc, ok := c.(SigninCommand)
	if !ok {
		return errors.New("SigninCommand: unexpected command")
	}
	return h.signin.Signin(ctx, sc.email, sc.password)
}
