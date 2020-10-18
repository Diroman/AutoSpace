package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"parking/internal/database"
	"parking/internal/predictor"
)

type Server struct {
	host      string
	port      int
	Predictor *predictor.Predictor
	Database  *database.Database
	Router    *mux.Router
}

func NewServer(host string, port int, db *database.Database, predictor *predictor.Predictor) *Server {
	r := mux.NewRouter()

	server := &Server{
		host:     host,
		port:     port,
		Router:   r,
		Database: db,
		Predictor: predictor,
	}

	r.Use(LoggingMiddleware)
	r.HandleFunc("/send-frame", server.SendFrame)
	r.HandleFunc("/login", server.Login)
	r.HandleFunc("/all-cameras", server.GetCameras)
	r.HandleFunc("/get-frame", server.GetFrame)
	r.HandleFunc("/get-free-space", server.GetFreeSpace)
	r.HandleFunc("/send-email", server.SendEmail)

	return server
}

func (s *Server) Run() {
	address := fmt.Sprintf("%s:%v", s.host, s.port)

	srv := &http.Server{
		Handler: s.Router,
		Addr:    address,
	}

	log.Printf("Server run on http://%s\n", address)

	log.Fatal(srv.ListenAndServe())
}
