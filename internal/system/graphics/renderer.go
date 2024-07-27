package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/entity"
)

// Renderer defines the interface for rendering systems.
type Renderer interface {
	Draw(*ebiten.Image, *entity.Entities)
}
