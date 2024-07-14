package graphics

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/entity"
)

type Drawer struct {
	op *ebiten.DrawImageOptions
}

func NewDrawer() *Drawer {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest

	return &Drawer{
		op: op,
	}
}

func (d *Drawer) Draw(screen *ebiten.Image, entities *entity.Entities) {
	for _, e := range *entities {
		d.drawEntity(screen, e)
	}
}

func (d *Drawer) drawEntity(screen *ebiten.Image, e *entity.Entity) {
	entityType, found := e.Type()
	if !found {
		return
	}

	bounds, found := e.Bounds()
	if !found {
		return
	}

	switch entityType {
	case entity.TypeBox:
		orange := color.RGBA{R: 255, G: 165, B: 0, A: 255}
		d.drawSolidColor(screen, bounds, orange)
	case entity.TypeGround:
		darkGrey := color.RGBA{R: 169, G: 169, B: 169, A: 255}
		d.drawSolidColor(screen, bounds, darkGrey)
	case entity.TypePlayer:
		white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
		d.drawSolidColor(screen, bounds, white)
	default:
		// Do nothing
	}

}

func (d *Drawer) drawSolidColor(screen *ebiten.Image, bounds image.Rectangle, c color.RGBA) {
	d.op.GeoM.Reset()
	d.op.GeoM.Translate(float64(bounds.Min.X), float64(bounds.Min.Y))

	img := ebiten.NewImage(bounds.Dx(), bounds.Dy())
	img.Fill(c)

	screen.DrawImage(img, d.op)
}
