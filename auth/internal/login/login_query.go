package login

import (
	"context"
	"errors"

	"github.com/apascualco/apascualco-auth/kit/query"
)

const LoginQueryType query.Type = "apascualco.1.query.login"

// LoginQuery command.Command interface
type LoginQuery struct {
	email    string
	password string
}

func NewLoginQuery(email, password string) LoginQuery {
	return LoginQuery{
		email:    email,
		password: password,
	}
}

func (l LoginQuery) Type() query.Type {
	return LoginQueryType
}

type LoginQueryHandler struct {
	login Login
}

func NewLoginQueryHandler(login Login) LoginQueryHandler {
	return LoginQueryHandler{login: login}
}

// Handle implements command.Handle interface
func (h LoginQueryHandler) Handle(ctx context.Context, cmd query.Query) (interface{}, error) {
	c, ok := cmd.(LoginQuery)
	if !ok {
		return "", errors.New("LoginQueryHandler: unexpected query")
	}

	return h.login.Login(ctx, c.email, c.password)
}
