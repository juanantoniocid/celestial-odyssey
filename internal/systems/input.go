package systems

import (
	"celestial-odyssey/internal/component"
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
	log.Println("Updating character based on input")
	input, inputOk := character.GetComponent(component.InputComponent).(*component.Input)
	velocity, velOk := character.GetComponent(component.VelocityComponent).(*component.Velocity)

	if inputOk && velOk {
		log.Println("Updating character based on input: ", input, velocity)
		velocity.X = 0
		if input.Left {
			log.Println("Move to left")
			velocity.X = -moveSpeed
		}
		if input.Right {
			log.Println("Move to right")
			velocity.X = moveSpeed
		}
		if input.Jump {
			log.Println("Let's jump")
			velocity.Y = jumpSpeed
			input.Jump = false // Reset jump after applying it
		}
	}

	is.applyPhysics(character)
}

func (is *InputHandler) applyPhysics(character *entity.GameEntity) {
	log.Println("Applying physics to character")
	velocity, ok := character.GetComponent(component.VelocityComponent).(*component.Velocity)
	position, posOk := character.GetComponent(component.PositionComponent).(*component.Position)

	if ok && posOk {
		log.Println("Applying physics to character: ", velocity, position)
		position.X += velocity.X
		position.Y += velocity.Y
	}
}
