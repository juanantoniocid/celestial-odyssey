package system

import (
	"celestial-odyssey/internal/component"
	"celestial-odyssey/internal/entity"
)

func entityIsSupported(e *entity.Entity, entities *entity.Entities) bool {
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

		if entityBounds.Max.Y == otherBounds.Min.Y &&
			(entityBounds.Min.X >= otherBounds.Min.X && entityBounds.Min.X <= otherBounds.Max.X ||
				entityBounds.Max.X <= otherBounds.Min.X && entityBounds.Max.X <= otherBounds.Max.X) {
			return true
		}
	}

	return false
}

func entityPositionAndVelocity(entity *entity.Entity) (component.Position, component.Velocity, bool) {
	position, positionFound := entity.Position()
	velocity, velocityFound := entity.Velocity()
	if !positionFound || !velocityFound {
		return component.Position{}, component.Velocity{}, false
	}

	return position, velocity, true
}

func entityCollidesOnTop(entity, other *entity.Entity) bool {
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

func entityCollidesOnBottom(entity, other *entity.Entity) bool {
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

func entityCollidesOnLeft(entity, other *entity.Entity) bool {
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

func entityCollidesOnRight(entity, other *entity.Entity) bool {
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
