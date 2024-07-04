package input

import (
	"celestial-odyssey/internal/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

// KeyboardHandler is responsible for handling keyboard input.
type KeyboardHandler struct{}

// NewKeyboardHandler creates a new KeyboardHandler instance.
func NewKeyboardHandler() *KeyboardHandler {
	return &KeyboardHandler{}
}

// UpdatePlayer updates the player based on the keyboard input.
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
