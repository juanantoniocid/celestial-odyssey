package graphics

import (
	"celestial-odyssey/world/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

type Renderer struct {
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) DrawPlayer(screen, playerImage *ebiten.Image, player *entities.Player) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest

	if player.Facing() == entities.Right {
		// Flip the image horizontally when facing left
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(playerImage.Bounds().Dx()), 0)
	}
	op.GeoM.Translate(float64(player.Position().X), float64(player.Position().Y))

	screen.DrawImage(playerImage, op)
}
