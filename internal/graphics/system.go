package graphics

import (
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
