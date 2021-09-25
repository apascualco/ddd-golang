// Package classification auth API Documentation
//
//	Schemes: http, https
//	Host: localhost:8080
//	BasePath: /
//	Version: 0.0.1
//	Contact:
//
//	Consume:
//	- application/json
//
//	Produces:
//	- application/json
//
//	swagger:meta
package server

import (
	"fmt"
	"log"
	"net/http"

	auth "github.com/apascualco/apascualco-auth/internal"
	"github.com/apascualco/apascualco-auth/internal/platform/server/handler/health"
	"github.com/apascualco/apascualco-auth/internal/platform/server/handler/login"
	"github.com/apascualco/apascualco-auth/internal/platform/server/handler/signin"
	"github.com/apascualco/apascualco-auth/internal/platform/server/middleware/logging"
	"github.com/apascualco/apascualco-auth/internal/platform/server/middleware/recovery"
	"github.com/apascualco/apascualco-auth/kit/command"
	"github.com/apascualco/apascualco-auth/kit/event"
	"github.com/apascualco/apascualco-auth/kit/query"
	"github.com/gin-gonic/gin"
	openApiMiddleware "github.com/go-openapi/runtime/middleware"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	userRepository auth.UserRepository
	secret         string

	commandBus command.Bus
	queryBus   query.Bus
	eventBus   event.Bus
}

func New(host string, port uint, commandBus command.Bus, queryBus query.Bus, eventBus event.Bus,
	userRepository auth.UserRepository, secret string) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		userRepository: userRepository,
		secret:         secret,

		commandBus: commandBus,
		queryBus:   queryBus,
		eventBus:   eventBus,
	}
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	Initialice(s.commandBus, s.queryBus, s.userRepository, s.secret, s.eventBus)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) RegisterRoutes() {
	v1 := s.engine.Group("/", func(c *gin.Context) {
		if "application/vnd.auth.v1+json" == c.GetHeader("accept") {
			c.Next()
		}
		c.AbortWithStatus(http.StatusNotFound)
	})
	v1.Use(recovery.Middleware(), logging.Middleware())
	v1.POST("/login", login.Login(s.queryBus))
	v1.POST("/signin", signin.Signin(s.commandBus))
	s.engine.GET("/health", health.CheckHandler())
}

func (s *Server) ConfigureSwagger() {
	o := openApiMiddleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sw := openApiMiddleware.SwaggerUI(o, nil)
	s.engine.GET("/docs", gin.WrapH(sw))
	s.engine.GET("/swagger.yaml", func(c *gin.Context) {
		c.File("./swagger.yaml")
	})
}
