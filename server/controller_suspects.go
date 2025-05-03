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

func addSuspectsRoutes(r chi.Router, llm llms.Model, db *gorm.DB) {
	r.Get("/suspects", func(w http.ResponseWriter, r *http.Request) {
		var suspects []schema.Suspect
		db.Find(&suspects)
		w.Write([]byte(fmt.Sprintf("%v", suspects)))
	})

	r.Get("/suspects/{id}", func(w http.ResponseWriter, r *http.Request) {
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
