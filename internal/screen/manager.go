package screen

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/world/entities"
)

type Manager struct {
	currentLevel Level
	renderer     Renderer
}

type Renderer interface {
	DrawPlayer(screen *ebiten.Image, player *entities.Player)
}

func NewManager(width, height int, player *entities.Player, renderer Renderer) *Manager {
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
