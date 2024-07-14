package systems

import "celestial-odyssey/internal/entity"

type System interface {
	Update(entityCollection *entity.Collection)
}

type Manager struct {
	systemHandlers []System
}

func NewManager(sh ...System) *Manager {
	systemHandler := make([]System, 0)

	for _, s := range sh {
		systemHandler = append(systemHandler, s)
	}

	return &Manager{
		systemHandlers: systemHandler,
	}
}

func (m *Manager) Update(entityCollection *entity.Collection) {
	for _, sh := range m.systemHandlers {
		sh.Update(entityCollection)
	}
}
