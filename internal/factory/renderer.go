package factory

import "celestial-odyssey/internal/system/graphics"

func CreateRenderer() *graphics.RenderManager {
	simpleRenderer := graphics.NewSimpleRenderer()
	rendererManager := graphics.NewRendererManager(simpleRenderer)

	return rendererManager
}
