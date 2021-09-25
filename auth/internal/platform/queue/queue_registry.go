package queue

import (
	"time"

	auth "github.com/apascualco/apascualco-auth/internal"
	"github.com/apascualco/apascualco-auth/internal/platform/queue/handler"
)

type QueueHandler struct {
	user     string
	password string
	port     string
	vhost    string

	userRepository auth.UserRepository
	wattingToStart int
}

func NewQueueHandler(u, p, port, v string, ur auth.UserRepository, w int) QueueHandler {

	return QueueHandler{
		user:     u,
		password: p,
		port:     port,
		vhost:    v,

		userRepository: ur,
		wattingToStart: w,
	}
}

func (j QueueHandler) InitializeHandlers() {
	if j.wattingToStart > 0 {
		time.Sleep(time.Duration(j.wattingToStart) * time.Second)
	}
	authHandler := handler.NewAuthHandler(j.userRepository)
	authHandlerQueue := NewRabbitMQ(j.user, j.password, j.port, j.vhost, authHandler.Queue())
	authHandlerQueue.ReadMessages(authHandler)
}
