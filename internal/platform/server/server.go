package server

type Server struct {
	port int32
}

func (s *Server) Start() {
	// addr := fmt.Sprintf(":%d", s.port)
}

func New(port int32) *Server {
	return &Server{port: port}
}
