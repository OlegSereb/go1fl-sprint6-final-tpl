package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Server представляет HTTP-сервер
type Server struct {
	logger *log.Logger
	server *http.Server
}

// NewServer создаёт и настраивает сервер
func NewServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger: logger,
		server: srv,
	}
}

// Start запускает сервер
func (s *Server) Start() error {
	s.logger.Printf("Сервер запущен на порту 8080")
	return s.server.ListenAndServe()
}
