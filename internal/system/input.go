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

func (is *Input) Update(entities *entity.Entities) {
	for _, e := range *entities {
		inputMap, found := e.InputMap()
		if !found {
			continue
		}

		is.updatePlayer(e, inputMap)
	}
}

func (is *Input) updatePlayer(e *entity.Entity, inputMap component.InputMap) {
	var action component.Action

	if ebiten.IsKeyPressed(inputMap.Left) {
		action.Left = true
	}

	if ebiten.IsKeyPressed(inputMap.Right) {
		action.Right = true
	}

	if ebiten.IsKeyPressed(inputMap.Jump) {
		action.Jump = true
	}

	e.SetAction(action)
}
