package screen

import (
	"celestial-odyssey/internal/graphics"
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/world/entities"
)

type Manager struct {
	currentLevel Level
	renderer     *graphics.Renderer
}

func NewScreenManager(width, height int, player *entities.Player) *Manager {
	renderer := graphics.NewRenderer()
	return &Manager{
		currentLevel: NewLevel1(width, height, player, renderer),
		renderer:     renderer,
	}
}

func (m *Manager) Update() {
	m.currentLevel.Update()
}

func (m *Manager) Draw(screen *ebiten.Image) {
	m.currentLevel.Draw(screen)
}
