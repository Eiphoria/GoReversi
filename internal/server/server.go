package server

import (
	"log/slog"
	"net/http"

	"github.com/Eiphoria/GoReversi/internal/service"
)

type Server struct {
	serv    http.Server
	service *service.Service
	logger  *slog.Logger
}

func New(service *service.Service, logger *slog.Logger) *Server {
	s := Server{
		service: service,
		logger:  logger,
	}
	m := http.NewServeMux()
	m.HandleFunc("/api/v1/health", s.health)
	m.HandleFunc("/api/v1/auth/reg", s.register)
	s.serv.Addr = "127.0.0.1:8080"
	s.serv.Handler = m
	return &s
}

func (s *Server) Run() error {

	return s.serv.ListenAndServe()
}
