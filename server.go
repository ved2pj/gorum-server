package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	engine *gin.Engine

	srv *http.Server
}

func NewServer() *Server {
	return &Server{
		engine: gin.Default(),
		srv: &http.Server{
			Addr: "127.0.0.1:8080",
		},
	}
}

func (server *Server) Start() error {
	routerGroup := server.engine.Group("/groum")
	server.setupApis(routerGroup)

	if err := server.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
		return err
	}

	return nil
}

func (server *Server) ShutDown(ctx context.Context) error {
	return server.srv.Shutdown(ctx)
}

func (server *Server) setupApis(group *gin.RouterGroup) {
	group.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gorum!")
	})
}
