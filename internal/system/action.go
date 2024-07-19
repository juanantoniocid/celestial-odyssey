package system

import (
	"celestial-odyssey/internal/entity"
)

type Action struct {
}

func NewAction() *Action {
	return &Action{}
}

const (
	moveSpeed = 2.0
	jumpSpeed = -10.0
)

func (m *Action) Update(entities *entity.Entities) {
	for _, e := range *entities {
		m.update(e, entities)
	}
}

func (m *Action) update(e *entity.Entity, entities *entity.Entities) {
	m.applyHorizontalMovement(e)
	m.applyVerticalMovement(e, entities)
}

func (m *Action) applyHorizontalMovement(e *entity.Entity) {
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

	e.SetVelocity(velocity)
}

func (m *Action) applyVerticalMovement(e *entity.Entity, entities *entity.Entities) {
	action, found := e.Action()
	if !found {
		return
	}

	velocity, found := e.Velocity()
	if !found {
		return
	}

	if action.Jump {
		if entityIsSupported(e, entities) {
			velocity.VY = jumpSpeed
		}
	}

	e.SetVelocity(velocity)
}
