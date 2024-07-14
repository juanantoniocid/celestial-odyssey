package system

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/entity"
)

// DrawSystem is an interface that defines the DrawSystem method.
type DrawSystem interface {
	Draw(*ebiten.Image, *entity.Entities)
}

// DrawSystems is a struct that holds a slice of DrawSystem.
type DrawSystems struct {
	draws []DrawSystem
}

// NewDrawSystems creates a new DrawSystems struct.
func NewDrawSystems(ds ...DrawSystem) *DrawSystems {
	drawSystem := make([]DrawSystem, 0)

	for _, d := range ds {
		drawSystem = append(drawSystem, d)
	}

	return &DrawSystems{
		draws: drawSystem,
	}
}

// Draw calls the DrawSystem method on each DrawSystem in the DrawSystems struct.
func (ds *DrawSystems) Draw(screen *ebiten.Image, entities *entity.Entities) {
	for _, d := range ds.draws {
		d.Draw(screen, entities)
	}
}

type SimpleDraw struct {
	op *ebiten.DrawImageOptions
}

func NewSimpleDraw() *SimpleDraw {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest

	return &SimpleDraw{
		op: op,
	}
}

func (sd *SimpleDraw) Draw(screen *ebiten.Image, entities *entity.Entities) {
	for _, e := range *entities {
		sd.drawEntity(screen, e)
	}
}

func (sd *SimpleDraw) drawEntity(screen *ebiten.Image, e *entity.Entity) {
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

func (sd *SimpleDraw) drawSolidColor(screen *ebiten.Image, bounds image.Rectangle, c color.RGBA) {
	sd.op.GeoM.Reset()
	sd.op.GeoM.Translate(float64(bounds.Min.X), float64(bounds.Min.Y))

	img := ebiten.NewImage(bounds.Dx(), bounds.Dy())
	img.Fill(c)

	screen.DrawImage(img, sd.op)
}
