package graphics

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Renderer struct {
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) DrawPlayer(screen, image *ebiten.Image, position image.Point) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest
	op.GeoM.Translate(float64(position.X), float64(position.Y))

	screen.DrawImage(image, op)
}
