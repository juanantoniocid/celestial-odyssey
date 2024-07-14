package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/entity"
)

// Draw is an interface that defines the Draw method.
type Draw interface {
	Draw(*ebiten.Image, *entity.Entities)
}

// DrawSystem is a struct that holds a slice of Draw.
type DrawSystem struct {
	draws []Draw
}

// NewDrawSystem creates a new DrawSystem struct.
func NewDrawSystem(ds ...Draw) *DrawSystem {
	drawSystem := make([]Draw, 0)

	for _, d := range ds {
		drawSystem = append(drawSystem, d)
	}

	return &DrawSystem{
		draws: drawSystem,
	}
}

// Draw calls the Draw method on each Draw in the DrawSystem struct.
func (ds *DrawSystem) Draw(screen *ebiten.Image, entities *entity.Entities) {
	for _, d := range ds.draws {
		d.Draw(screen, entities)
	}
}
