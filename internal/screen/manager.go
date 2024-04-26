package screen

import (
	"celestial-odyssey/world/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

type Manager struct {
	CurrentScreen *Level1
}

func NewScreenManager(width, height int, player *entities.Player) *Manager {
	return &Manager{
		CurrentScreen: NewGameScreen(width, height, player),
	}
}

func (m *Manager) Update() {
	m.CurrentScreen.Update()
}

func (m *Manager) Draw(screen *ebiten.Image) {
	m.CurrentScreen.Draw(screen)
}
