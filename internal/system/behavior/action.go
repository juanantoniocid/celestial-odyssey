package behavior

import (
	"celestial-odyssey/internal/entity"
	"celestial-odyssey/internal/system/util"
)

// Action defines a behavior system that handles horizontal and vertical movement for entities.
type Action struct {
	moveSpeed float64
	jumpSpeed float64
}

// NewAction creates a new instance of the Action system.
func NewAction(moveSpeed, jumpSpeed float64) *Action {
	return &Action{
		moveSpeed: moveSpeed,
		jumpSpeed: jumpSpeed,
	}
}

// Update updates the state of the entities to apply movements based on actions.
func (a *Action) Update(entities *entity.Entities) {
	for _, e := range *entities {
		a.update(e, entities)
	}
}

func (a *Action) update(e *entity.Entity, entities *entity.Entities) {
	a.applyHorizontalMovement(e)
	a.applyVerticalMovement(e, entities)
}

func (a *Action) applyHorizontalMovement(e *entity.Entity) {
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
		velocity.VX = -a.moveSpeed
	}
	if action.Right {
		velocity.VX = a.moveSpeed
	}

	e.SetVelocity(velocity)
}

func (a *Action) applyVerticalMovement(e *entity.Entity, entities *entity.Entities) {
	action, found := e.Action()
	if !found {
		return
	}

	velocity, found := e.Velocity()
	if !found {
		return
	}

	if action.Jump && util.IsEntityGrounded(e, entities) {
		velocity.VY = a.jumpSpeed
	}

	e.SetVelocity(velocity)
}
