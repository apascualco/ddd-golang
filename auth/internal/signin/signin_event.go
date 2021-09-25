package signin

import (
	"encoding/json"
	"time"

	"github.com/apascualco/apascualco-auth/kit/event"
)

const SigninEventType event.Type = "apascualco.auth.1.event.auth.new"

type SigninEvent struct {
	event.BaseEvent
	uuid     string
	email    string
	password string
}

func NewSigninEvent(uuid, email, password string) SigninEvent {
	return SigninEvent{
		uuid:     uuid,
		email:    email,
		password: password,

		BaseEvent: event.NewBaseEvent(uuid),
	}
}

func (e SigninEvent) Type() event.Type {
	return SigninEventType
}

func (e SigninEvent) UUID() string {
	return e.uuid
}

func (e SigninEvent) Email() string {
	return e.email
}

type signinEventJson struct {
	UUID        string    `json:"uuid"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	AggregateID string    `json:"aggregateID"`
	OccurredOn  time.Time `json:"occurredOn"`
}

func (e SigninEvent) Marshaller() ([]byte, error) {
	le := signinEventJson{
		UUID:        e.uuid,
		Email:       e.email,
		Password:    e.password,
		AggregateID: e.AggregateID(),
		OccurredOn:  e.OccurredOn(),
	}
	return json.Marshal(le)
}
