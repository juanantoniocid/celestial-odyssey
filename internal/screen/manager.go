package screen

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/world/entities"
)

type Manager struct {
	CurrentLevel Level
}

func NewScreenManager(width, height int, player *entities.Player) *Manager {
	return &Manager{
		CurrentLevel: NewLevel1(width, height, player),
	}
}

func (m *Manager) Update() {
	m.CurrentLevel.Update()
}

func (m *Manager) Draw(screen *ebiten.Image) {
	m.CurrentLevel.Draw(screen)
}
