package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"

	"celestial-odyssey/internal/entity"
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
	characterInput, err := character.Input()
	if err != nil {
		log.Println("failed to get character input")
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		log.Println("left")
		characterInput.Left = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		log.Println("right")
		characterInput.Right = true
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		log.Println("jump")
		characterInput.Jump = true
	}

	character.SetInput(characterInput)
}
