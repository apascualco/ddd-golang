package auth

//go:generate mockgen -source=internal/event_handler.go -destination internal/platform/queue/mocks/mock_event_handler.go
type EventHandler interface {
	EventHandler(body string) error
}
