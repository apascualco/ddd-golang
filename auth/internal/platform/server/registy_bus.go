package server

import (
	auth "github.com/apascualco/apascualco-auth/internal"
	"github.com/apascualco/apascualco-auth/internal/login"

	"github.com/apascualco/apascualco-auth/internal/signin"
	"github.com/apascualco/apascualco-auth/kit/command"
	"github.com/apascualco/apascualco-auth/kit/event"
	"github.com/apascualco/apascualco-auth/kit/query"
)

func Initialice(commandBus command.Bus, queryBus query.Bus, ur auth.UserRepository,
	secret string, eventBus event.Bus) {
	l := login.NewLogin(ur, eventBus, secret)
	loginHandler := login.NewLoginQueryHandler(l)
	queryBus.Register(login.LoginQueryType, loginHandler)

	s := signin.NewSignin(ur, eventBus)
	signinHandler := signin.NewSigninCommandHandler(s)
	commandBus.Register(signin.SiginCommandType, signinHandler)
}
