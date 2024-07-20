package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/entity"
)

// Renderer is an interface that defines the Draw method for rendering systems.
type Renderer interface {
	Draw(*ebiten.Image, *entity.Entities)
}
