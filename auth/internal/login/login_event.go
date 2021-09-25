package login

import (
	"encoding/json"
	"time"

	"github.com/apascualco/apascualco-auth/kit/event"
)

const LoginEventType event.Type = "apascualco.auth.1.event.auth.login"

type LoginEvent struct {
	event.BaseEvent
	uuid  string
	email string
}

func NewLoginEvent(uuid, email string) LoginEvent {
	return LoginEvent{
		uuid:  uuid,
		email: email,

		BaseEvent: event.NewBaseEvent(uuid),
	}
}

func (e LoginEvent) Type() event.Type {
	return LoginEventType
}

func (e LoginEvent) UUID() string {
	return e.uuid
}

func (e LoginEvent) Email() string {
	return e.email
}

type loginEventJson struct {
	UUID        string    `json:"uuid"`
	Email       string    `json:"email"`
	AggregateID string    `json:"aggregateID"`
	OccurredOn  time.Time `json:"occurredOn"`
}

func (e LoginEvent) Marshaller() ([]byte, error) {
	le := loginEventJson{
		UUID:        e.uuid,
		Email:       e.email,
		AggregateID: e.AggregateID(),
		OccurredOn:  e.OccurredOn(),
	}
	return json.Marshal(le)
}
