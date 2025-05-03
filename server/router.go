package server

import (
	"github.com/AnthonyThomasson/go-scratchpad/agent"
	"github.com/AnthonyThomasson/go-scratchpad/schema"
	"github.com/go-chi/chi"
	"github.com/tmc/langchaingo/llms"
	"gorm.io/gorm"
)

func getRouter(llm llms.Model, db *gorm.DB) *chi.Mux {
	// delete all murders
	db.Delete(&schema.Murder{})
	db.Delete(&schema.Clue{})
	db.Delete(&schema.Suspect{})
	db.Delete(&schema.Location{})

	murder := agent.NewMurder(
		llm,
		"A figure lies sprawled across a cracked concrete floor, a gun tossed nearby.",
		schema.Clue{
			Name:        "Bar reciept from 'The Red Room'",
			Description: "A bar reciept found at the scene, in the victim's pocket.",
		},
	)

	db.Create(murder)

	r := chi.NewRouter()
	addSuspectsRoutes(r, llm, db)
	addMurderRoutes(r, llm, db)
	addCluesRoutes(r, llm, db)
	return r
}
