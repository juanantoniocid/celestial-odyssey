package systems

import (
	"celestial-odyssey/internal/debug"
	"celestial-odyssey/internal/entity"
	"log"
)

type InputHandler struct {
}

func NewInputHandler() *InputHandler {
	return &InputHandler{}
}

const (
	moveSpeed = 2.0
	jumpSpeed = -5.0
)

func (is *InputHandler) Update(character *entity.GameEntity) {
	input, err := character.Input()
	if err != nil {
		debug.Log("Failed to get input component: %e", err)
		return
	}

	velocity, err := character.Velocity()
	if err != nil {
		debug.Log("Failed to get velocity component: %e", err)
		return
	}

	velocity.VX = 0
	if input.Left {
		log.Println("Move to left")
		velocity.VX = -moveSpeed
	}
	if input.Right {
		log.Println("Move to right")
		velocity.VX = moveSpeed
	}
	if input.Jump {
		log.Println("Let's jump")
		velocity.VY = jumpSpeed
		input.Jump = false // Reset jump after applying it
	}

	is.applyPhysics(character)
}

func (is *InputHandler) applyPhysics(character *entity.GameEntity) {
	velocity, err := character.Velocity()
	if err != nil {
		debug.Log("Failed to get velocity component: %e", err)
		return
	}

	position, err := character.Position()
	if err != nil {
		debug.Log("Failed to get position component: %e", err)
		return
	}

	position.X += velocity.VX
	position.Y += velocity.VY
}
