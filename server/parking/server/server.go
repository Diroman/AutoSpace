package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"parking/handlers"
)

type Server struct {
	host      string
	port      int
	Predictor string
	Router    *mux.Router
}

func NewServer(host string, port int) *Server {
	r := mux.NewRouter()
	r.Use(handlers.LoggingMiddleware)

	r.HandleFunc("/send-frame", handlers.SendFrame)
	r.HandleFunc("/login", handlers.Login)

	return &Server{
		host:   host,
		port:   port,
		Router: r,
	}
}

func (s *Server) Run() {
	address := fmt.Sprintf("%s:%v", s.host, s.port)

	srv := &http.Server{
		Handler: s.Router,
		Addr: address,
	}

	log.Printf("Server run on http://%s\n", address)

	log.Fatal(srv.ListenAndServe())
}
