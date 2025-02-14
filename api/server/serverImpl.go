package server

import (
	"fmt"
	"go-todo/handlers"
	"log"
	"net/http"
)

type HTTPServer struct {
	Port   uint
	server *http.Server
}

func NewHTTPServer(port uint) *HTTPServer {
	r := http.NewServeMux()

	r.HandleFunc("POST /profile", handlers.CreateProfileHandler)

	return &HTTPServer{
		Port: port,
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port), // Convert port to :<port> string
			Handler: r,
		},
	}
}

func (s *HTTPServer) Start() error {
	log.Printf("Server starting on port %d\n", s.Port)
	return s.server.ListenAndServe()
}
