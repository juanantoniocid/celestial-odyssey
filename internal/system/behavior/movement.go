package behavior

import (
	"celestial-odyssey/internal/entity"
	"celestial-odyssey/internal/system/util"
)

// Movement defines a system for handling entity movement and collisions.
type Movement struct{}

// NewMovement creates a new instance of the Movement system.
func NewMovement() *Movement {
	return &Movement{}
}

// Update updates the position of each entity in the system based on their velocity.
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
	position, velocity, found := util.EntityPositionAndVelocity(entity)
	if !found {
		return
	}

	position.Y += velocity.VY
	entity.SetPosition(position)

	m.handleVerticalCollisions(entity, entities)
	m.handleHorizontalCollisions(entity, entities)
}

func (m *Movement) handleVerticalCollisions(entity *entity.Entity, entities *entity.Entities) {
	for _, other := range *entities {
		if entity == other {
			continue
		}

		m.handleVerticalCollision(entity, other)
	}
}

func (m *Movement) handleVerticalCollision(entity, other *entity.Entity) {
	m.handleBottomCollision(entity, other)
	m.handleTopCollision(entity, other)
}

func (m *Movement) handleBottomCollision(entity, other *entity.Entity) {
	entityPosition, found := entity.Position()
	if !found {
		return
	}

	entitySize, found := entity.Size()
	if !found {
		return
	}

	entityVelocity, found := entity.Velocity()
	if !found {
		return
	}

	otherPosition, found := other.Position()
	if !found {
		return
	}

	if util.EntityCollidesOnBottom(entity, other) {
		// Align the bottom of the entity with the top of the other entity
		entityPosition.Y = otherPosition.Y - entitySize.Height
		entity.SetPosition(entityPosition)

		// Stop the entity from moving vertically
		entityVelocity.VY = 0
		entity.SetVelocity(entityVelocity)
	}
}

func (m *Movement) handleTopCollision(entity, other *entity.Entity) {
	entityPosition, found := entity.Position()
	if !found {
		return
	}

	entityVelocity, found := entity.Velocity()
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

	if util.EntityCollidesOnTop(entity, other) {
		// Align the top of the entity with the bottom of the other entity
		entityPosition.Y = otherPosition.Y + otherSize.Height
		entity.SetPosition(entityPosition)

		// Stop the entity from moving vertically
		entityVelocity.VY = 0
		entity.SetVelocity(entityVelocity)
	}
}

func (m *Movement) updateHorizontalMovement(entity *entity.Entity, entities *entity.Entities) {
	position, velocity, found := util.EntityPositionAndVelocity(entity)
	if !found {
		return
	}

	position.X += velocity.VX
	entity.SetPosition(position)

	m.handleHorizontalCollisions(entity, entities)
	m.handleVerticalCollisions(entity, entities)
}

func (m *Movement) handleHorizontalCollisions(entity *entity.Entity, entities *entity.Entities) {
	for _, other := range *entities {
		if entity == other {
			continue
		}

		m.handleHorizontalCollision(entity, other)
	}
}

func (m *Movement) handleHorizontalCollision(entity, other *entity.Entity) {
	m.handleLeftCollision(entity, other)
	m.handleRightCollision(entity, other)
}

func (m *Movement) handleLeftCollision(entity, other *entity.Entity) {
	entityPosition, found := entity.Position()
	if !found {
		return
	}

	entityVelocity, found := entity.Velocity()
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

	if util.EntityCollidesOnLeft(entity, other) {
		// Align the left side of the entity with the right side of the other entity
		entityPosition.X = otherPosition.X + otherSize.Width
		entity.SetPosition(entityPosition)

		// Stop the entity from moving horizontally
		entityVelocity.VX = 0
		entity.SetVelocity(entityVelocity)
	}
}

func (m *Movement) handleRightCollision(entity, other *entity.Entity) {
	entityPosition, found := entity.Position()
	if !found {
		return
	}

	entitySize, found := entity.Size()
	if !found {
		return
	}

	entityVelocity, found := entity.Velocity()
	if !found {
		return
	}

	otherPosition, found := other.Position()
	if !found {
		return
	}

	if util.EntityCollidesOnRight(entity, other) {
		// Align the right side of the entity with the left side of the other entity
		entityPosition.X = otherPosition.X - entitySize.Width
		entity.SetPosition(entityPosition)

		// Stop the entity from moving horizontally
		entityVelocity.VX = 0
		entity.SetVelocity(entityVelocity)
	}
}
