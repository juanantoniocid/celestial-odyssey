package system

import (
	"celestial-odyssey/internal/entity"
)

type Movement struct {
}

func NewMovement() *Movement {
	return &Movement{}
}

const (
	moveSpeed = 2.0
	jumpSpeed = -5.0
)

func (m *Movement) Update(entities *entity.Entities) {
	for _, e := range *entities {
		m.update(e)
	}
}

func (m *Movement) update(e *entity.Entity) {
	action, found := e.Action()
	if !found {
		return
	}

	velocity, found := e.Velocity()
	if !found {
		return
	}

	velocity.VX = 0
	if action.Left {
		velocity.VX = -moveSpeed
	}
	if action.Right {
		velocity.VX = moveSpeed
	}
	if action.Jump {
		velocity.VY = jumpSpeed
		action.Jump = false // Reset jump after applying it
	}

	e.SetVelocity(velocity)

	m.applyPhysics(e)
}

func (m *Movement) applyPhysics(character *entity.Entity) {
	velocity, found := character.Velocity()
	if !found {
		return
	}

	position, found := character.Position()
	if !found {
		return
	}

	position.X += velocity.VX
	position.Y += velocity.VY

	character.SetPosition(position)
}
