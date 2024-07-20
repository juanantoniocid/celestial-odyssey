package behavior

import "celestial-odyssey/internal/entity"

// UpdateSystemManager is a struct that holds a slice of UpdateSystem.
type UpdateSystemManager struct {
	updateSystems []UpdateSystem
}

// NewUpdateSystemManager creates a new UpdateSystemManager struct.
func NewUpdateSystemManager(ss ...UpdateSystem) *UpdateSystemManager {
	updateSystems := make([]UpdateSystem, 0)
	updateSystems = append(updateSystems, ss...)

	return &UpdateSystemManager{
		updateSystems: updateSystems,
	}
}

// Update calls the Update method on each UpdateSystem in the UpdateSystemManager struct.
func (us *UpdateSystemManager) Update(entities *entity.Entities) {
	for _, u := range us.updateSystems {
		u.Update(entities)
	}
}
