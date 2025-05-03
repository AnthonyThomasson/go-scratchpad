package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AnthonyThomasson/go-scratchpad/schema"
	"github.com/go-chi/chi"
	"github.com/tmc/langchaingo/llms"
	"gorm.io/gorm"
)

func addMurderRoutes(r chi.Router, llm llms.Model, db *gorm.DB) {
	r.Get("/murders", func(w http.ResponseWriter, r *http.Request) {
		var murder schema.Murder
		db.First(&murder)
		w.Write([]byte(fmt.Sprintf("%v", murder)))
	})

	r.Get("/murders/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		var murder schema.Murder
		db.First(&murder, id)
		w.Write([]byte(fmt.Sprintf("%v", murder)))
	})
}
