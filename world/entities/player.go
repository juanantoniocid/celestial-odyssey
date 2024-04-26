package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	X, Y  float64
	Speed float64
	Image *ebiten.Image
}

func NewPlayer(x, y, speed float64, image *ebiten.Image) *Player {
	return &Player{
		X:     x,
		Y:     y,
		Speed: speed,
		Image: image,
	}
}

func (p *Player) Update() {
	// Update player position based on input
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.X -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.X += 2
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest
	op.GeoM.Translate(p.X, p.Y)

	screen.DrawImage(p.Image, op)
}
