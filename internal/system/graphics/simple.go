package graphics

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/component"
	"celestial-odyssey/internal/entity"
)

// SimpleRenderer is a basic implementation of the Renderer interface.
type SimpleRenderer struct {
	op *ebiten.DrawImageOptions
}

// NewSimpleRenderer creates a new instance of SimpleRenderer.
func NewSimpleRenderer() *SimpleRenderer {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest

	return &SimpleRenderer{
		op: op,
	}
}

// Draw renders the entities to the screen.
func (sr *SimpleRenderer) Draw(screen *ebiten.Image, entities *entity.Entities) {
	for _, e := range *entities {
		sr.drawEntity(screen, e)
	}
}

func (sr *SimpleRenderer) drawEntity(screen *ebiten.Image, e *entity.Entity) {
	bounds, found := e.Bounds()
	if !found {
		return
	}

	sprite, found := e.Sprite()
	if found {
		sr.drawSprite(screen, bounds, sprite)
		return
	}

	color, found := e.Color()
	if found {
		sr.drawSolidColor(screen, bounds, color)
	}
}

func (sr *SimpleRenderer) drawSolidColor(screen *ebiten.Image, bounds image.Rectangle, c component.Color) {
	sr.op.GeoM.Reset()
	sr.op.GeoM.Translate(float64(bounds.Min.X), float64(bounds.Min.Y))

	img := ebiten.NewImage(bounds.Dx(), bounds.Dy())
	img.Fill(c.Color)

	screen.DrawImage(img, sr.op)
}

func (sr *SimpleRenderer) drawSprite(screen *ebiten.Image, bounds image.Rectangle, sprite component.Sprite) {
	sr.op.GeoM.Reset()
	sr.op.GeoM.Translate(float64(bounds.Min.X), float64(bounds.Min.Y))

	screen.DrawImage(sprite.Image, sr.op)
}
