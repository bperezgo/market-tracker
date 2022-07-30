package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"markettracker.com/pkg/command"
	"markettracker.com/tracker/internal/platform/server/handler"
)

type Server struct {
	httpAddr        string
	shutdownTimeout time.Duration
	engine          *gin.Engine

	commandBus command.Bus
}

func New(host string, port int32, cmdBus command.Bus) *Server {
	addr := fmt.Sprintf("%s:%d", host, port)
	srv := &Server{
		httpAddr:        addr,
		shutdownTimeout: 10 * time.Second,
	}
	srv.registerRoutes()
	return srv
}

func (s *Server) registerRoutes() {
	// TODO: Middlewares to implement: RecoveryMiddleware, LoggingMiddleware
	s.engine.GET("/health", handler.Health)
	s.engine.POST("/bvc-asset", handler.BvcAsset(s.commandBus))
}

func (s *Server) Start(ctx context.Context) error {
	log.Println("Server running on", s.httpAddr)
	srv := http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()
	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}
