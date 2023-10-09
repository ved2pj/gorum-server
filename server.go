package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	srv *http.Server

	engine *gin.Engine
}

func NewServer() *Server {
	engine := gin.Default()
	return &Server{
		engine: engine,
		srv: &http.Server{
			Addr:    ":8080",
			Handler: engine,
		},
	}
}

func (server *Server) Start() error {
	group := server.engine.Group("/gorum")
	server.setupApis(group)

	if err := server.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
