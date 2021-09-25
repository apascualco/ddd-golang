package query

import "context"

type Bus interface {
	Dispatch(context.Context, Query) (interface{}, error)
	Register(Type, Handler)
}

//go:generate mockgen -source=kit/query/query.go -destination kit/query/mockquery/mock_query.go
type Type string

type Query interface {
	Type() Type
}

type Handler interface {
	Handle(context.Context, Query) (interface{}, error)
}
