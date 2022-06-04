package server

import (
	"fmt"
	"log"

	"markettracker.com/app"
)

type Server struct {
	port int32
}

var server *Server

func (s *Server) Start() {
	addr := fmt.Sprintf(":%d", s.port)
	a := app.GetInstance()
	a.Start()
	if err := a.Run(addr); err != nil {
		log.Panic("App crash")
	}
}

func InitServer(port int32) *Server {
	if server != nil {
		return server
	}
	return &Server{port: port}
}
