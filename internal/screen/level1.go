package screen

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/world/entities"
)

type Level1 struct {
	player        *entities.Player
	Width, Height int
}

func NewLevel1(width, height int, player *entities.Player) *Level1 {
	player.MoveToRightBoundary(width)
	player.MoveToBottomBoundary(height)
	player.SetSpeed(1)

	return &Level1{
		player: player,
		Width:  width,
		Height: height,
	}
}

func (l1 *Level1) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		l1.player.MoveLeft()
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		l1.player.MoveRight()
	}

	l1.player.Update()

	// Constrain player's movement within the screen bounds:
	if l1.player.Left() < 0 {
		l1.player.MoveToLeftBoundary(0)
	} else if l1.player.Right() > l1.Width {
		l1.player.MoveToRightBoundary(l1.Width)
	}
}

func (l1 *Level1) Draw(screen *ebiten.Image) {
	l1.player.Draw(screen)
}
