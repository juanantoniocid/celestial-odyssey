package behavior

import "celestial-odyssey/internal/entity"

// UpdateSystem is an interface that defines the Update method for updating updateSystems.
type UpdateSystem interface {
	Update(*entity.Entities)
}
