package event

import (
	"context"

	"github.com/apascualco/apascualco-auth/kit/event"
)

// InMemoryEventBus is an in-memory implementation of the event.Bus.
type InMemoryEventBus struct {
	handlers map[event.Type][]event.Handler
}

func NewInMemoryEventBus() *InMemoryEventBus {
	return &InMemoryEventBus{
		handlers: make(map[event.Type][]event.Handler),
	}
}

// Publish implements the event.Bus interface.
func (b *InMemoryEventBus) Publish(ctx context.Context, events []event.Event) error {
	for _, evt := range events {
		handlers, ok := b.handlers[evt.Type()]
		if !ok {
			return nil
		}

		for _, handler := range handlers {
			handler.Handle(ctx, evt)
		}
	}

	return nil
}
