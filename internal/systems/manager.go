package systems

import "celestial-odyssey/internal/entity"

type System interface {
	Update(entityCollection *entity.Collection)
}

type Manager struct {
	inputSystem     System
	physicsSystem   System
	movementSystem  System
	collisionSystem System
}

func NewManager(inputSystem System, physicsSystem System, movementSystem System, collisionSystem System) *Manager {
	return &Manager{
		inputSystem:     inputSystem,
		physicsSystem:   physicsSystem,
		movementSystem:  movementSystem,
		collisionSystem: collisionSystem,
	}
}

func (m *Manager) Update(entityCollection *entity.Collection) {
	m.inputSystem.Update(entityCollection)
	m.physicsSystem.Update(entityCollection)
	m.movementSystem.Update(entityCollection)
	m.collisionSystem.Update(entityCollection)
}
