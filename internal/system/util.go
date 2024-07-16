package system

import (
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
