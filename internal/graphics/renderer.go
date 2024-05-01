package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/world/entities"
)

type Renderer struct {
	playerImage *ebiten.Image
}

func NewRenderer(playerImage *ebiten.Image) *Renderer {
	return &Renderer{
		playerImage: playerImage,
	}
}

func (r *Renderer) DrawPlayer(screen *ebiten.Image, player *entities.Player) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest

	if player.Facing() == entities.Right {
		// Flip the image horizontally when facing left
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(r.playerImage.Bounds().Dx()), 0)
	}
	op.GeoM.Translate(float64(player.Position().X), float64(player.Position().Y))

	screen.DrawImage(r.playerImage, op)
}
