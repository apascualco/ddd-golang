package command

import (
	"context"
	"errors"
	"testing"

	"github.com/apascualco/apascualco-auth/kit/command"
	"github.com/stretchr/testify/assert"
)

type HandlerTest struct {
	err error
}

func (ht HandlerTest) Handle(ctx context.Context, q command.Command) error {
	return ht.err
}

func (ht HandlerTest) Type() command.Type {
	return "test"
}

func TestCommand(t *testing.T) {

	t.Run("Given a correct command bus should not return err", func(t *testing.T) {
		// Given
		ht := HandlerTest{}

		// When
		commandBus := NewCommandBus()
		commandBus.Register(ht.Type(), ht)

		// Then
		err := commandBus.Dispatch(context.Background(), ht)
		assert.NoError(t, err)

	})

	t.Run("Given a correct command bus with a custom error", func(t *testing.T) {
		// Given
		var CustomError = errors.New("Custom error")
		ht := HandlerTest{err: CustomError}

		// When
		commandBus := NewCommandBus()
		commandBus.Register(ht.Type(), ht)

		// Then
		err := commandBus.Dispatch(context.Background(), ht)
		assert.ErrorIs(t, err, CustomError)
	})

	t.Run("Given and unregister command bus should return error", func(t *testing.T) {
		// Given
		ht := HandlerTest{}

		// When
		commandBus := NewCommandBus()

		// Then
		err := commandBus.Dispatch(context.Background(), ht)
		assert.ErrorIs(t, err, ErrorCommandBus)
	})
}
