package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/entity"
)

// RenderManager manages a collection of renderers.
type RenderManager struct {
	renderers []Renderer
}

// NewRendererManager creates a new instance of RenderManager.
func NewRendererManager(rs ...Renderer) *RenderManager {
	renderers := make([]Renderer, 0)
	renderers = append(renderers, rs...)

	return &RenderManager{
		renderers: renderers,
	}
}

// Draw calls the Draw method on each renderer in the manager.
func (rm *RenderManager) Draw(screen *ebiten.Image, entities *entity.Entities) {
	for _, r := range rm.renderers {
		r.Draw(screen, entities)
	}
}
