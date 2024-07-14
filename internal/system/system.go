package system

import "celestial-odyssey/internal/entity"

// System is an interface that defines the Update method.
type System interface {
	Update(*entity.Entities)
}

// Systems is a struct that holds a slice of System.
type Systems struct {
	systems []System
}

// NewSystems creates a new Systems struct.
func NewSystems(ss ...System) *Systems {
	systems := make([]System, 0)

	for _, s := range ss {
		systems = append(systems, s)
	}

	return &Systems{
		systems: systems,
	}
}

// Update calls the Update method on each System in the Systems struct.
func (m *Systems) Update(entities *entity.Entities) {
	for _, sh := range m.systems {
		sh.Update(entities)
	}
}
