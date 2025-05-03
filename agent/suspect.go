package agent

import "github.com/AnthonyThomasson/go-scratchpad/schema"

// NewSuspect creates a new suspect with the given profile and location
func NewSuspect(id int, name string, age int, location schema.Location) *schema.Suspect {
	return &schema.Suspect{
		ID:       id,
		Guilty:   false,
		Name:     name,
		Age:      age,
		Location: location,
	}
}
