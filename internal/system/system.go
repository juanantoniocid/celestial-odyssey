package system

import "celestial-odyssey/internal/entity"

// System is an interface that defines the Update method.
type System interface {
	Update(*entity.Collection)
}

// Systems is a struct that holds a slice of System.
type Systems struct {
	systems []System
}

// NewSystems creates a new Systems struct.
func NewSystems(sh ...System) *Systems {
	systemHandler := make([]System, 0)

	for _, s := range sh {
		systemHandler = append(systemHandler, s)
	}

	return &Systems{
		systems: systemHandler,
	}
}

// Update calls the Update method on each System in the Systems struct.
func (m *Systems) Update(entityCollection *entity.Collection) {
	for _, sh := range m.systems {
		sh.Update(entityCollection)
	}
}
