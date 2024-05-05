package screen

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/util"
	"celestial-odyssey/world/entities"
)

type Manager struct {
	currentLevel Level
	renderer     Renderer
}

type Renderer interface {
	DrawPlayer(screen *ebiten.Image, player *entities.Player)
}

func NewManager(dimensions util.Dimensions, player *entities.Player, renderer Renderer) *Manager {
	return &Manager{
		currentLevel: NewLevel1(dimensions, player, renderer),
		renderer:     renderer,
	}
}

func (m *Manager) Update() {
	m.currentLevel.Update()
}

func (m *Manager) Draw(screen *ebiten.Image) {
	m.currentLevel.Draw(screen)
}
