package event

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Bus interface {
	Publish(context.Context, []Event) error
}

// Handler defines the expected behaviour from an event handler.
type Handler interface {
	Handle(context.Context, Event) error
}

type Type string

type Event interface {
	ID() string
	AggregateID() string
	OccurredOn() time.Time
	Type() Type
	Marshaller() ([]byte, error)
}

type BaseEvent struct {
	eventID     string
	aggregateID string
	occurredOn  time.Time
}

func NewBaseEvent(aggregateID string) BaseEvent {
	return BaseEvent{
		eventID:     uuid.New().String(),
		aggregateID: aggregateID,
		occurredOn:  time.Now(),
	}
}

func (b BaseEvent) ID() string {
	return b.eventID
}

func (b BaseEvent) OccurredOn() time.Time {
	return b.occurredOn
}

func (b BaseEvent) AggregateID() string {
	return b.aggregateID
}
