package system

import (
	"log"

	"celestial-odyssey/internal/entity"
)

type Movement struct {
}

func NewMovement() *Movement {
	return &Movement{}
}

func (m *Movement) Update(entities *entity.Entities) {
	for _, e := range *entities {
		m.update(e, entities)
	}
}

func (m *Movement) update(e *entity.Entity, entities *entity.Entities) {
	m.updateVerticalMovement(e, entities)
	m.updateHorizontalMovement(e, entities)
}

func (m *Movement) updateVerticalMovement(entity *entity.Entity, entities *entity.Entities) {
	position, velocity, found := entityPositionAndVelocity(entity)
	if !found {
		return
	}

	position.Y += velocity.VY
	entity.SetPosition(position)

	m.handleVerticalCollisions(entity, entities)
	m.handleHorizontalCollisions(entity, entities)
}

func (m *Movement) updateHorizontalMovement(entity *entity.Entity, entities *entity.Entities) {
	position, velocity, found := entityPositionAndVelocity(entity)
	if !found {
		return
	}

	position.X += velocity.VX
	entity.SetPosition(position)

	m.handleHorizontalCollisions(entity, entities)
	m.handleVerticalCollisions(entity, entities)
}

func (m *Movement) handleVerticalCollisions(entity *entity.Entity, entities *entity.Entities) {
	for _, other := range *entities {
		if entity == other {
			continue
		}

		m.handleVerticalCollision(entity, other)
	}
}

func (m *Movement) handleHorizontalCollisions(entity *entity.Entity, entities *entity.Entities) {
	for _, other := range *entities {
		if entity == other {
			continue
		}

		m.handleHorizontalCollision(entity, other)
	}
}

func (m *Movement) handleVerticalCollision(entity, other *entity.Entity) {
	entityPosition, found := entity.Position()
	if !found {
		return
	}

	entitySize, found := entity.Size()
	if !found {
		return
	}

	otherPosition, found := other.Position()
	if !found {
		return
	}

	otherSize, found := other.Size()
	if !found {
		return
	}

	entityVelocity, found := entity.Velocity()
	if !found {
		return
	}

	if entityCollidesOnBottom(entity, other) {
		log.Println("Collides on its ground", entityPosition, entitySize, otherPosition, otherSize)
		entityPosition.Y = otherPosition.Y - entitySize.Height
		entity.SetPosition(entityPosition)
		entityVelocity.VY = 0
		entity.SetVelocity(entityVelocity)
		log.Println("Ground collision fixed", entityPosition, entitySize, otherPosition, otherSize)
	}

	if entityCollidesOnTop(entity, other) {
		log.Println("Collides on its top", entityPosition, entitySize, otherPosition, otherSize)
		entityPosition.Y = otherPosition.Y + otherSize.Height
		entity.SetPosition(entityPosition)
		entityVelocity.VY = 0
		entity.SetVelocity(entityVelocity)
		log.Println("Top collision fixed", entityPosition, entitySize, otherPosition, otherSize)
	}
}

func (m *Movement) handleHorizontalCollision(entity, other *entity.Entity) {
	entityPosition, found := entity.Position()
	if !found {
		return
	}

	entitySize, found := entity.Size()
	if !found {
		return
	}

	otherPosition, found := other.Position()
	if !found {
		return
	}

	otherSize, found := other.Size()
	if !found {
		return
	}

	if entityCollidesOnLeft(entity, other) {
		log.Println("Collides on its left", entityPosition, entitySize, otherPosition, otherSize)
		entityPosition.X = otherPosition.X + otherSize.Width
		entity.SetPosition(entityPosition)
		log.Println("Left collision fixed", entityPosition, entitySize, otherPosition, otherSize)
	}

	if entityCollidesOnRight(entity, other) {
		log.Println("Collides on its right", entityPosition, entitySize, otherPosition, otherSize)
		entityPosition.X = otherPosition.X - entitySize.Width
		entity.SetPosition(entityPosition)
		log.Println("Right collision fixed", entityPosition, entitySize, otherPosition, otherSize)
	}
}
