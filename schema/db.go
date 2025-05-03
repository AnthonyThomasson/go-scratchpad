package schema

import (
	"github.com/tmc/langchaingo/llms"
	"gorm.io/gorm"
)

type Clue struct {
	gorm.Model
	ID          int
	Name        string
	Description string

	Murder Murder `gorm:"foreignKey:MurderID"`

	MurderID uint
}

type Location struct {
	gorm.Model
	ID          int
	Name        string
	Description string

	Murder Murder `gorm:"foreignKey:MurderID"`

	MurderID uint
}

type Murder struct {
	gorm.Model
	ID      int
	Details string

	Clues    []Clue
	Suspects []Suspect
}

// MessageSuspect messages a suspect with the given message
func (m *Murder) MessageSuspect(model llms.Model, id int, message string) string {
	return "I have no brain"
}

type Suspect struct {
	gorm.Model
	ID     int
	Guilty bool
	Name   string
	Age    int

	Location Location `gorm:"foreignKey:LocationID"`
	Murder   Murder   `gorm:"foreignKey:MurderID"`

	LocationID uint
	MurderID   uint
}

// SetGuilty marks the suspect as guilty
func (s *Suspect) SetGuilty() *Suspect {
	s.Guilty = true
	return s
}
