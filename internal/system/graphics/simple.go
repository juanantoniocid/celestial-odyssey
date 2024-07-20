package graphics

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/entity"
)

// SimpleRenderer is a basic implementation of the Renderer interface.
type SimpleRenderer struct {
	op *ebiten.DrawImageOptions
}

// NewSimpleRenderer creates a new SimpleRenderer struct.
func NewSimpleRenderer() *SimpleRenderer {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest

	return &SimpleRenderer{
		op: op,
	}
}

// Draw renders the entities to the screen.
func (sd *SimpleRenderer) Draw(screen *ebiten.Image, entities *entity.Entities) {
	for _, e := range *entities {
		sd.drawEntity(screen, e)
	}
}

func (sd *SimpleRenderer) drawEntity(screen *ebiten.Image, e *entity.Entity) {
	bounds, found := e.Bounds()
	if !found {
		return
	}

	entityColor, found := e.Color()
	if !found {
		return
	}

	sd.drawSolidColor(screen, bounds, entityColor.Color)
}

func (sd *SimpleRenderer) drawSolidColor(screen *ebiten.Image, bounds image.Rectangle, c color.RGBA) {
	sd.op.GeoM.Reset()
	sd.op.GeoM.Translate(float64(bounds.Min.X), float64(bounds.Min.Y))

	img := ebiten.NewImage(bounds.Dx(), bounds.Dy())
	img.Fill(c)

	screen.DrawImage(img, sd.op)
}
