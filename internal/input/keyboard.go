package input

import (
	"celestial-odyssey/internal/world/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

type KeyboardHandler struct{}

func NewKeyboardHandler() *KeyboardHandler {
	return &KeyboardHandler{}
}

func (kh *KeyboardHandler) UpdatePlayer(player *entities.Player) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		player.MoveLeft()
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		player.MoveRight()
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		player.Jump()
	}
}
