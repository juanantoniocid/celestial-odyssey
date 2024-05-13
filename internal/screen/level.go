package screen

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/world/entities"
)

type Level interface {
	Update()
	Draw(screen *ebiten.Image)
}

type LevelImpl struct {
	player   *entities.Player
	scenario *ebiten.Image
	renderer Renderer
}

func NewLevel(cfg config.Screen, player *entities.Player, renderer Renderer) *LevelImpl {
	levelWidth := cfg.Dimensions.Width
	levelHeight := cfg.Dimensions.Height
	groundLeft := image.Point{X: 0, Y: levelHeight - 1}

	player.SetPlayArea(image.Rect(0, 0, levelWidth-1, levelHeight-1))
	player.SetPositionAtBottomLeft(groundLeft)
	player.SetSpeed(2)

	background := loadBackgroundImage("assets/images/landing-site.png")

	return &LevelImpl{
		player:   player,
		scenario: background,
		renderer: renderer,
	}
}

func loadBackgroundImage(file string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(file)
	if err != nil {
		log.Fatal("failed to load scenario image:", err)
		return nil
	}

	return img
}

func (l1 *LevelImpl) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		l1.player.MoveLeft()
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		l1.player.MoveRight()
	}

	l1.player.Update()
}

func (l1 *LevelImpl) Draw(screen *ebiten.Image) {
	l1.drawScenario(screen)
	l1.renderer.DrawPlayer(screen, l1.player)
}

func (l1 *LevelImpl) drawScenario(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest

	screen.DrawImage(l1.scenario, op)
}
