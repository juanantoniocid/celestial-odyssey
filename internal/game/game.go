package game

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/screen"
	"celestial-odyssey/world/entities"
)

type Game struct {
	ScreenManager *screen.Manager
}

func NewGame(player *entities.Player) *Game {
	screenManager := screen.NewScreenManager(320, 200, player)

	return &Game{
		ScreenManager: screenManager,
	}
}

func (g *Game) Update() error {
	g.ScreenManager.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ScreenManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 200 // Return the fixed game screen size
}
