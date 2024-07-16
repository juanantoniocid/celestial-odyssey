package system

import "celestial-odyssey/internal/entity"

type Movement struct {
}

func NewMovement() *Movement {
	return &Movement{}
}

func (m *Movement) Update(entities *entity.Entities) {
	for _, e := range *entities {
		m.update(e)
	}
}

func (m *Movement) update(e *entity.Entity) {
	velocity, found := e.Velocity()
	if !found {
		return
	}

	position, found := e.Position()
	if !found {
		return
	}

	position.X += velocity.VX
	position.Y += velocity.VY

	e.SetPosition(position)
}
