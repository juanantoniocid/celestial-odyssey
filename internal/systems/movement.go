package systems

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

func (is *Movement) Update(entityCollection *entity.Collection) {
	for _, e := range *entityCollection {
		is.update(e)
	}
}

func (is *Movement) update(e *entity.GameEntity) {
	input, found := e.Input()
	if !found {
		return
	}

	velocity, found := e.Velocity()
	if !found {
		return
	}

	velocity.VX = 0
	if input.Left {
		velocity.VX = -moveSpeed
	}
	if input.Right {
		velocity.VX = moveSpeed
	}
	if input.Jump {
		velocity.VY = jumpSpeed
		input.Jump = false // Reset jump after applying it
	}

	e.SetVelocity(velocity)

	is.applyPhysics(e)
}

func (is *Movement) applyPhysics(character *entity.GameEntity) {
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
