package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ScreenManager interface {
	Init()
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
}

type Game struct {
	screenManager ScreenManager
}

func NewGame(screenManager ScreenManager) *Game {
	return &Game{
		screenManager: screenManager,
	}
}

func (g *Game) Init() {
	g.screenManager.Init()
}

func (g *Game) Update() error {
	return g.screenManager.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.screenManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.screenManager.Layout(outsideWidth, outsideHeight)
}
