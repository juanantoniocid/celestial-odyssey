package behavior

import "celestial-odyssey/internal/entity"

// UpdateSystemManager manages a collection of update systems.
type UpdateSystemManager struct {
	updateSystems []UpdateSystem
}

// NewUpdateSystemManager creates a new instance of UpdateSystemManager.
func NewUpdateSystemManager(ss ...UpdateSystem) *UpdateSystemManager {
	updateSystems := make([]UpdateSystem, 0)
	updateSystems = append(updateSystems, ss...)

	return &UpdateSystemManager{
		updateSystems: updateSystems,
	}
}

// Update calls the Update method on each update system in the manager.
func (us *UpdateSystemManager) Update(entities *entity.Entities) {
	for _, u := range us.updateSystems {
		u.Update(entities)
	}
}
