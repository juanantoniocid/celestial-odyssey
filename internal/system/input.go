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

func (is *Input) Update(entityCollection *entity.Collection) {
	for _, e := range *entityCollection {
		inputMap, found := e.InputMap()
		if !found {
			continue
		}

		is.updatePlayer(e, inputMap)
	}
}

func (is *Input) updatePlayer(e *entity.Entity, inputMap component.InputMap) {
	var input component.Input

	if ebiten.IsKeyPressed(inputMap.Left) {
		input.Left = true
	}

	if ebiten.IsKeyPressed(inputMap.Right) {
		input.Right = true
	}

	if ebiten.IsKeyPressed(inputMap.Jump) {
		input.Jump = true
	}

	e.SetInput(input)
}
