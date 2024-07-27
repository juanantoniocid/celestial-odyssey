package util

import (
	"celestial-odyssey/internal/component"
	"celestial-odyssey/internal/entity"
)

// IsEntityGrounded checks if an entity is grounded by checking its bounds against other entities.
func IsEntityGrounded(e *entity.Entity, entities *entity.Entities) bool {
	entityBounds, entityBoundsFound := e.Bounds()
	if !entityBoundsFound {
		return false
	}

	for _, other := range *entities {
		if e == other {
			continue
		}

		otherBounds, otherBoundsFound := other.Bounds()
		if !otherBoundsFound {
			continue
		}

		// Check if the bottom of the entity is aligned with the top of the other entity
		if entityBounds.Max.Y == otherBounds.Min.Y {
			// Check if there is horizontal overlap
			if (entityBounds.Min.X >= otherBounds.Min.X && entityBounds.Min.X < otherBounds.Max.X) ||
				(entityBounds.Max.X > otherBounds.Min.X && entityBounds.Max.X <= otherBounds.Max.X) ||
				(entityBounds.Min.X <= otherBounds.Min.X && entityBounds.Max.X >= otherBounds.Max.X) {
				return true
			}
		}
	}

	return false
}

// EntityPositionAndVelocity returns the position and velocity of an entity if both are found.
func EntityPositionAndVelocity(entity *entity.Entity) (component.Position, component.Velocity, bool) {
	position, positionFound := entity.Position()
	velocity, velocityFound := entity.Velocity()
	if !positionFound || !velocityFound {
		return component.Position{}, component.Velocity{}, false
	}

	return position, velocity, true
}

// EntityCollidesOnTop checks if an entity collides on the top with another entity.
func EntityCollidesOnTop(entity, other *entity.Entity) bool {
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

// EntityCollidesOnBottom checks if an entity collides on the bottom with another entity.
func EntityCollidesOnBottom(entity, other *entity.Entity) bool {
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

	return entityBounds.Max.Y > otherBounds.Min.Y && entityBounds.Min.Y < otherBounds.Min.Y
}

// EntityCollidesOnLeft checks if an entity collides on the left with another entity.
func EntityCollidesOnLeft(entity, other *entity.Entity) bool {
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

// EntityCollidesOnRight checks if an entity collides on the right with another entity.
func EntityCollidesOnRight(entity, other *entity.Entity) bool {
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
