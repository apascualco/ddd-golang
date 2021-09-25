package command

import (
	"context"
	"errors"

	"github.com/apascualco/apascualco-auth/kit/command"
)

var ErrorCommandBus = errors.New("CommandBus: Command bus handler not found")

type CommandBus struct {
	handlers map[command.Type]command.Handler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[command.Type]command.Handler),
	}
}

func (b *CommandBus) Dispatch(ctx context.Context, cmd command.Command) error {
	handler, ok := b.handlers[cmd.Type()]
	if !ok {
		return ErrorCommandBus
	}
	errChannel := make(chan error)
	go func() {
		err := handler.Handle(ctx, cmd)
		if err != nil {
			errChannel <- err
			return
		}
		errChannel <- nil
	}()

	return <-errChannel
}

func (b *CommandBus) Register(cmdType command.Type, handler command.Handler) {
	b.handlers[cmdType] = handler
}
