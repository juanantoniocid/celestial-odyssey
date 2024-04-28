package game

import (
	"celestial-odyssey/internal/screen"
	"celestial-odyssey/world/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	screenManager *screen.Manager
	width         int
	height        int
}

func NewGame(gameWidth, gameHeight int) *Game {
	player := entities.NewPlayer()

	screenManager := screen.NewScreenManager(gameWidth, gameHeight, player)

	return &Game{
		screenManager: screenManager,
		width:         gameWidth,
		height:        gameHeight,
	}
}

func (g *Game) Update() error {
	g.screenManager.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.screenManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}
