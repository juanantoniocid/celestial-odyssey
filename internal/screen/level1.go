package screen

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/world/entities"
)

type Level1 struct {
	player        *entities.Player
	Width, Height int
}

func NewLevel1(width, height int, player *entities.Player) *Level1 {
	groundLeft := image.Point{X: 0, Y: height}
	player.SetPositionAtBottomLeft(groundLeft)
	player.SetSpeed(2)

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
}

func (l1 *Level1) Draw(screen *ebiten.Image) {
	l1.player.Draw(screen)
}
