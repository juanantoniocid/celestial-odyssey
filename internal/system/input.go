package system

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/component"
	"celestial-odyssey/internal/entity"
)

type Input struct {
}

func NewInput() *Input {
	return &Input{}
}

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
