package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ScreenManager interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Game struct {
	screenManager ScreenManager
	width         int
	height        int
}

func NewGame(gameWidth, gameHeight int, screenManager ScreenManager) *Game {
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
