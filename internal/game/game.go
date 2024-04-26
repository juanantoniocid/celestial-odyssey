package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"celestial-odyssey/internal/screen"
	"celestial-odyssey/world/entities"
)

type Game struct {
	screenManager *screen.Manager
	width         int
	height        int
}

func NewGame(gameWidth, gameHeight int) *Game {
	playerImage := loadPlayerImage()
	player := entities.NewPlayer(100, 100, 0, playerImage)

	screenManager := screen.NewScreenManager(gameWidth, gameHeight, player)

	return &Game{
		screenManager: screenManager,
		width:         gameWidth,
		height:        gameHeight,
	}
}

func loadPlayerImage() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile("assets/images/player.png")
	if err != nil {
		log.Fatal(err)
	}

	return img
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
