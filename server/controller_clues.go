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

func addCluesRoutes(r chi.Router, llm llms.Model, db *gorm.DB) {
	r.Get("/clues", func(w http.ResponseWriter, r *http.Request) {
		var clues []schema.Clue
		db.Find(&clues)
		w.Write([]byte(fmt.Sprintf("%v", clues)))
	})

	r.Get("/clues/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		var suspect schema.Suspect
		db.First(&suspect, id)
		w.Write([]byte(fmt.Sprintf("%v", suspect)))
	})
}
