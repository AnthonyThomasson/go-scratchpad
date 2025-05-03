package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tmc/langchaingo/llms"
	"gorm.io/gorm"
)

// Server is the main server struct
type Server struct {
	router  *chi.Mux
	address string
	llm     llms.Model
	db      *gorm.DB
}

// NewServer creates a new server with the given LLM
func NewServer(llm llms.Model, db *gorm.DB) *Server {
	return &Server{
		router:  getRouter(llm, db),
		address: ":8080",
		llm:     llm,
		db:      db,
	}
}

// Start starts the server
func (s *Server) Start() {
	http.ListenAndServe(s.address, s.router)
}
