package entities

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	X, Y  float64
	Image *ebiten.Image
}

func NewPlayer() *Player {
	// Load your player image here
	img, _, err := ebitenutil.NewImageFromFile("assets/images/player.png")
	if err != nil {
		log.Fatal(err)
	}
	return &Player{
		X:     100,
		Y:     100,
		Image: img,
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
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.Y += 2
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(p.Image, op)
}
