package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/entity"
)

// RenderManager is a struct that holds a slice of Renderer.
type RenderManager struct {
	renderers []Renderer
}

// NewRendererManager creates a new RenderManager struct.
func NewRendererManager(rs ...Renderer) *RenderManager {
	renderers := make([]Renderer, 0)
	renderers = append(renderers, rs...)

	return &RenderManager{
		renderers: renderers,
	}
}

// Draw calls the Renderer method on each Renderer in the RenderManager struct.
func (rm *RenderManager) Draw(screen *ebiten.Image, entities *entity.Entities) {
	for _, r := range rm.renderers {
		r.Draw(screen, entities)
	}
}
