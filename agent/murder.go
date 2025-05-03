package agent

import (
	"github.com/AnthonyThomasson/go-scratchpad/schema"
	"github.com/tmc/langchaingo/llms"
)

// NewMurder creates a new murder with the given details and clue
func NewMurder(llm llms.Model, details string, clue schema.Clue) *schema.Murder {
	return &schema.Murder{
		Details: details,
		Clues:   []schema.Clue{clue},
		Suspects: []schema.Suspect{
			*NewSuspect(
				1,
				"John Doe",
				30,
				schema.Location{
					Name:        "A bar named 'The Red Room'",
					Description: "A bar in the center of town, known for its dark atmosphere and live music.",
				},
			),
			*NewSuspect(
				2,
				"Peter Parker",
				25,
				schema.Location{
					Name:        "A coffee shop named 'The Bean'",
					Description: "A coffee shop nessled next to a park. It is very popular with the local college students.",
				},
			),
			*NewSuspect(
				3,
				"Mary Jane",
				22,
				schema.Location{
					Name:        "A flower shop named 'Blossom'",
					Description: "A flower shop that sells flowers for weddings and other special occasions.",
				},
			),
		},
	}
}
