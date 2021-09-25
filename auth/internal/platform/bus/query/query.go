package query

import (
	"context"
	"errors"

	"github.com/apascualco/apascualco-auth/kit/query"
)

var ErrorQueryBus = errors.New("QueryBus: Query bus handler not found")

type QueryBus struct {
	handlers map[query.Type]query.Handler
}

func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[query.Type]query.Handler),
	}
}

func (b *QueryBus) Dispatch(ctx context.Context, q query.Query) (interface{}, error) {
	handler, ok := b.handlers[q.Type()]
	if !ok {
		return nil, ErrorQueryBus
	}

	result, err := handler.Handle(ctx, q)
	return result, err
}

func (b *QueryBus) Register(queryType query.Type, handler query.Handler) {
	b.handlers[queryType] = handler
}
