package behavior

import (
	"celestial-odyssey/internal/entity"
	"celestial-odyssey/internal/system/util"
)

type Gravity struct {
	gravity float64
}

func NewGravity(gravity float64) *Gravity {
	return &Gravity{
		gravity: gravity,
	}
}

func (g *Gravity) Update(entities *entity.Entities) {
	for _, e := range *entities {
		g.update(e, entities)
	}
}

func (g *Gravity) update(e *entity.Entity, entities *entity.Entities) {
	action, found := e.Action()
	if !found {
		return
	}

	velocity, found := e.Velocity()
	if !found {
		return
	}

	if util.IsEntityGrounded(e, entities) {
		if !action.Jump {
			velocity.VY = 0
		}
	} else {
		velocity.VY += g.gravity
	}

	e.SetVelocity(velocity)
}
