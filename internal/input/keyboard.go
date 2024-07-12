package input

import (
	"celestial-odyssey/internal/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

// KeyboardHandler is responsible for handling keyboard input.
type KeyboardHandler struct{}

// NewKeyboardHandler creates a new KeyboardHandler instance.
func NewKeyboardHandler() *KeyboardHandler {
	return &KeyboardHandler{}
}

// UpdatePlayer updates the player based on the keyboard input.
func (kh *KeyboardHandler) UpdatePlayer(player *entity.Player) {
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

// UpdateCharacter updates the character based on the keyboard input.
func (kh *KeyboardHandler) UpdateCharacter(character *entity.GameEntity) {
	characterInput, found := character.Input()
	if !found {
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		characterInput.Left = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		characterInput.Right = true
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		characterInput.Jump = true
	}

	character.SetInput(characterInput)
}
