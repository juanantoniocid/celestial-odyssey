package game

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/world/entities"
)

type Game struct {
	// Add game state variables here
	player *entities.Player
}

func NewGame() *Game {
	return &Game{
		player: entities.NewPlayer(),
	}
}

func (g *Game) Update() error {
	g.player.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 200 // Return the fixed game screen size
}
