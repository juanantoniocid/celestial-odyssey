package behavior

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/component"
	"celestial-odyssey/internal/entity"
)

// Input defines a system that maps user inputs to entity actions.
type Input struct{}

// NewInput creates a new instance of the Input system.
func NewInput() *Input {
	return &Input{}
}

// Update maps user inputs to actions for each entity in the system.
func (i *Input) Update(entities *entity.Entities) {
	for _, e := range *entities {
		input, found := e.Input()
		if !found {
			continue
		}

		i.updatePlayer(e, input)
	}
}

func (i *Input) updatePlayer(e *entity.Entity, input component.Input) {
	var currentAction component.Action

	if ebiten.IsKeyPressed(input.Left) {
		currentAction.Left = true
	}

	if ebiten.IsKeyPressed(input.Right) {
		currentAction.Right = true
	}

	if ebiten.IsKeyPressed(input.Jump) {
		currentAction.Jump = true
	}

	e.SetAction(currentAction)
}
