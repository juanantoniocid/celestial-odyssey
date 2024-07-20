package factory

import "celestial-odyssey/internal/system/behavior"

func CreateUpdateSystem() *behavior.UpdateSystemManager {
	inputSystem := behavior.NewInput()
	actionSystem := behavior.NewAction()
	gravitySystem := behavior.NewGravity()
	movementSystem := behavior.NewMovement()

	return behavior.NewUpdateSystemManager(inputSystem, actionSystem, gravitySystem, movementSystem)
}
