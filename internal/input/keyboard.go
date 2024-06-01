package input

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/world/entities"
)

type KeyboardHandler struct {
	player *entities.Player
}

func NewKeyboardHandler(player *entities.Player) *KeyboardHandler {
	return &KeyboardHandler{
		player: player,
	}
}

func (kh *KeyboardHandler) UpdatePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		kh.player.MoveLeft()
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		kh.player.MoveRight()
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		kh.player.Jump()
	}
}
