package screen

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/world/entities"
)

type Level1 struct {
	player        *entities.Player
	Width, Height int
}

func NewGameScreen(width, height int, player *entities.Player) *Level1 {
	return &Level1{
		player: player,
		Width:  width,
		Height: height,
	}
}

func (l1 *Level1) Update() {
	l1.player.Update()

	// Constrain player's movement within the screen bounds:
	if l1.player.X < 0 {
		l1.player.X = 0
	} else if l1.player.X+float64(l1.player.Image.Bounds().Dx()) > float64(l1.Width) {
		l1.player.X = float64(l1.Width) - float64(l1.player.Image.Bounds().Dx())
	}

	if l1.player.Y < 0 {
		l1.player.Y = 0
	} else if l1.player.Y+float64(l1.player.Image.Bounds().Dy()) > float64(l1.Height) {
		l1.player.Y = float64(l1.Height) - float64(l1.player.Image.Bounds().Dy())
	}
}

func (l1 *Level1) Draw(screen *ebiten.Image) {
	l1.player.Draw(screen)
}
