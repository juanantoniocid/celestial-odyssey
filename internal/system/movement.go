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
	velocity, found := e.Velocity()
	if !found {
		return
	}

	position, found := e.Position()
	if !found {
		return
	}

	position.Y += velocity.VY
	e.SetPosition(position)
	m.handleVerticalCollisions(e, entities)
	m.handleHorizontalCollisions(e, entities)

	velocity, _ = e.Velocity()
	position, _ = e.Position()

	position.X += velocity.VX
	e.SetPosition(position)
	m.handleHorizontalCollisions(e, entities)
	m.handleVerticalCollisions(e, entities)
}

func (m *Movement) handleVerticalCollisions(entity *entity.Entity, entities *entity.Entities) {
	entityBounds, found := entity.Bounds()
	if !found {
		return
	}

	for _, other := range *entities {
		if entity == other {
			continue
		}

		otherBounds, otherFound := other.Bounds()
		if !otherFound {
			continue
		}

		if entityBounds.Overlaps(otherBounds) {
			m.handleVerticalCollision(entity, other)
		}
	}
}

func (m *Movement) handleHorizontalCollisions(entity *entity.Entity, entities *entity.Entities) {
	entityBounds, found := entity.Bounds()
	if !found {
		return
	}

	for _, other := range *entities {
		if entity == other {
			continue
		}

		otherBounds, otherFound := other.Bounds()
		if !otherFound {
			continue
		}

		if entityBounds.Overlaps(otherBounds) {
			m.handleHorizontalCollision(entity, other)
		}
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

	if m.entityCollidesOnItsGround(entity, other) {
		log.Println("Collides on its ground", entityPosition, entitySize, otherPosition, otherSize)
		entityPosition.Y = otherPosition.Y - entitySize.Height
		entity.SetPosition(entityPosition)
		entityVelocity.VY = 0
		entity.SetVelocity(entityVelocity)
		log.Println("Ground collision fixed", entityPosition, entitySize, otherPosition, otherSize)
	}

	if m.entityCollidesOnItsTop(entity, other) {
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

	if m.entityCollidesOnItsLeft(entity, other) {
		log.Println("Collides on its left", entityPosition, entitySize, otherPosition, otherSize)
		entityPosition.X = otherPosition.X + otherSize.Width
		entity.SetPosition(entityPosition)
		log.Println("Left collision fixed", entityPosition, entitySize, otherPosition, otherSize)
	}

	if m.entityCollidesOnItsRight(entity, other) {
		log.Println("Collides on its right", entityPosition, entitySize, otherPosition, otherSize)
		entityPosition.X = otherPosition.X - entitySize.Width
		entity.SetPosition(entityPosition)
		log.Println("Right collision fixed", entityPosition, entitySize, otherPosition, otherSize)
	}
}

func (m *Movement) entityCollidesOnItsGround(entity, other *entity.Entity) bool {
	entityBounds, found := entity.Bounds()
	if !found {
		return false
	}

	otherBounds, found := other.Bounds()
	if !found {
		return false
	}

	if !entityBounds.Overlaps(otherBounds) {
		return false
	}

	return entityBounds.Min.Y < otherBounds.Min.Y && entityBounds.Max.Y > otherBounds.Min.Y
}

func (m *Movement) entityCollidesOnItsTop(entity, other *entity.Entity) bool {
	entityBounds, found := entity.Bounds()
	if !found {
		return false
	}

	otherBounds, found := other.Bounds()
	if !found {
		return false
	}

	if !entityBounds.Overlaps(otherBounds) {
		return false
	}

	return entityBounds.Min.Y > otherBounds.Min.Y && entityBounds.Min.Y < otherBounds.Max.Y
}

func (m *Movement) entityCollidesOnItsLeft(entity, other *entity.Entity) bool {
	entityBounds, found := entity.Bounds()
	if !found {
		return false
	}

	otherBounds, found := other.Bounds()
	if !found {
		return false
	}

	if !entityBounds.Overlaps(otherBounds) {
		return false
	}

	return entityBounds.Min.X > otherBounds.Min.X && entityBounds.Min.X < otherBounds.Max.X
}

func (m *Movement) entityCollidesOnItsRight(entity, other *entity.Entity) bool {
	entityBounds, found := entity.Bounds()
	if !found {
		return false
	}

	otherBounds, found := other.Bounds()
	if !found {
		return false
	}

	if !entityBounds.Overlaps(otherBounds) {
		return false
	}

	return entityBounds.Max.X > otherBounds.Min.X && entityBounds.Max.X < otherBounds.Max.X
}
