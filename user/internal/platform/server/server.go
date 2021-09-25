// Package classification user API Documentation
//
//	Schemes: http, https
//	Host: localhost:8081
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

	"github.com/apascualco/apascualco-user/internal/platform/server/handler/health"
	"github.com/apascualco/apascualco-user/internal/platform/server/middleware/logging"
	"github.com/apascualco/apascualco-user/internal/platform/server/middleware/recovery"
	"github.com/gin-gonic/gin"
	openApiMiddleware "github.com/go-openapi/runtime/middleware"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
}

func New(host string, port uint) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
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
