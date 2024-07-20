package behavior

import (
	"celestial-odyssey/internal/entity"
	"celestial-odyssey/internal/system/util"
)

const (
	gravity = 0.5
)

type Gravity struct {
}

func NewGravity() *Gravity {
	return &Gravity{}
}

func (p *Gravity) Update(entities *entity.Entities) {
	for _, e := range *entities {
		p.update(e, entities)
	}
}

func (p *Gravity) update(e *entity.Entity, entities *entity.Entities) {
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
		velocity.VY += gravity
	}

	e.SetVelocity(velocity)
}
