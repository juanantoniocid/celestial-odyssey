package input

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/world/entities"
)

type KeyboardHandler struct{}

func NewKeyboardHandler() *KeyboardHandler {
	return &KeyboardHandler{}
}

func (ih *KeyboardHandler) Update(player *entities.Player) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		player.MoveLeft()
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		player.MoveRight()
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		player.Jump()
	}

	player.Update()
}
