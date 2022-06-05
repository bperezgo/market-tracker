package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpAddr        string
	shutdownTimeout time.Duration
}

func New(host string, port int32) *Server {
	addr := fmt.Sprintf("%s:%d", host, port)
	return &Server{
		httpAddr:        addr,
		shutdownTimeout: 10 * time.Second,
	}
}

func (s *Server) Start(ctx context.Context) error {
	log.Println("Server running on", s.httpAddr)
	srv := http.Server{
		Addr: s.httpAddr,
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
