package game

import (
	"celestial-odyssey/world/entities"
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/screen"
)

type Game struct {
	// Add game state variables here
	GameScreen *screen.Level1
}

func NewGame(player *entities.Player) *Game {
	gameScreen := screen.NewGameScreen(320, 200, player)

	return &Game{
		GameScreen: gameScreen,
	}
}

func (g *Game) Update() error {
	g.GameScreen.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.GameScreen.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 200 // Return the fixed game screen size
}
