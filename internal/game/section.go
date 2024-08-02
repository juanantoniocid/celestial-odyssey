package game

import (
	"celestial-odyssey/internal/entity"
)

// SectionID represents the ID of a section.
type SectionID string

// BasicSection represents a basic section in the game.
type BasicSection struct {
	entities *entity.Entities
}

// NewBasicSection creates a new basic section.
func NewBasicSection(entities *entity.Entities) *BasicSection {
	return &BasicSection{
		entities: entities,
	}
}

// Entities returns the entities in the section.
func (s *BasicSection) Entities() *entity.Entities {
	return s.entities
}
