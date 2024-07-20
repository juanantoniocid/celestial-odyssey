package factory

import "celestial-odyssey/internal/system/behavior"

const (
	moveSpeed = 2.0
	jumpSpeed = -10.0
	gravity   = 0.5
)

func CreateUpdateSystem() *behavior.UpdateSystemManager {
	inputSystem := behavior.NewInput()
	actionSystem := behavior.NewAction(moveSpeed, jumpSpeed)
	gravitySystem := behavior.NewGravity(gravity)
	movementSystem := behavior.NewMovement()

	return behavior.NewUpdateSystemManager(inputSystem, actionSystem, gravitySystem, movementSystem)
}
