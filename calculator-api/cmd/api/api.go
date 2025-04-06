package api

import (
	"fmt"
	service "mirzaadr/calculator-api/services"
	"net/http"
)

type APIServer struct {
	addr string
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	handler := service.NewHandler()
	handler.RegisterRoutes(router)

	fmt.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}
