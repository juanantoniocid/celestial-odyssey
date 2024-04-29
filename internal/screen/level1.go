package screen

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"celestial-odyssey/internal/graphics"
	"celestial-odyssey/world/entities"
)

type Level1 struct {
	player        *entities.Player
	Width, Height int
	renderer      *graphics.Renderer

	playerImage *ebiten.Image
}

func NewLevel1(width, height int, player *entities.Player, renderer *graphics.Renderer) *Level1 {
	playerImage := loadPlayerImage()
	groundLeft := image.Point{X: 0, Y: height}

	player.SetPlayArea(image.Rect(0, 0, width, height))
	player.SetDimensions(playerImage.Bounds().Dx(), playerImage.Bounds().Dy())
	player.SetPositionAtBottomLeft(groundLeft)
	player.SetSpeed(2)

	return &Level1{
		player:      player,
		Width:       width,
		Height:      height,
		renderer:    renderer,
		playerImage: playerImage,
	}
}

func loadPlayerImage() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile("assets/images/player.png")
	if err != nil {
		log.Fatal(err)
	}

	return img
}

func (l1 *Level1) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		l1.player.MoveLeft()
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		l1.player.MoveRight()
	}

	l1.player.Update()
}

func (l1 *Level1) Draw(screen *ebiten.Image) {
	l1.renderer.DrawPlayer(screen, l1.playerImage, l1.player)
}
