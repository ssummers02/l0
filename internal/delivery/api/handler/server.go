package handler

import (
	"context"
	"l0/internal/domain/service"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Server struct {
	httpServer *http.Server
	r          *mux.Router
	v          *validator.Validate
	services   *service.Service
}

//nolint:gomnd
func NewServer(port string, services *service.Service) *Server {
	r := mux.NewRouter()

	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + port,
			Handler:        r,
			MaxHeaderBytes: 1 << 20, // 1 MB
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
		r:        r,
		v:        validator.New(),
		services: services,
	}
}

func (s *Server) Run() error {
	s.initRoutes()

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) initRoutes() {
	router := s.r.PathPrefix("/api").
		Subrouter()

	router.HandleFunc("/orders/{id}", s.getOrders).
		Methods(http.MethodGet)
}
