package screen

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/graphics"
	"celestial-odyssey/world/entities"
)

const (
	sideMargin   = 8
	bottomMargin = 5
)

type ScenarioImpl struct {
	player     *entities.Player
	background *ebiten.Image
	renderer   *graphics.Renderer

	width  int
	height int
}

func NewScenario(player *entities.Player, background *ebiten.Image, renderer *graphics.Renderer, width, height int) *ScenarioImpl {
	return &ScenarioImpl{
		player:     player,
		background: background,
		renderer:   renderer,
		width:      width,
		height:     height,
	}
}

func (s *ScenarioImpl) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		s.player.MoveLeft()
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		s.player.MoveRight()
	}

	s.player.Update()

	return nil
}

func (s *ScenarioImpl) Draw(screen *ebiten.Image) {
	s.renderer.DrawBackground(screen, s.background)
	s.renderer.DrawPlayer(screen, s.player)
}

func (s *ScenarioImpl) ShouldTransitionRight() bool {
	return s.player.Position().X+s.player.Width() >= s.width
}

func (s *ScenarioImpl) ShouldTransitionLeft() bool {
	return s.player.Position().X <= 0
}

func (s *ScenarioImpl) SetPlayerPositionAtLeft() {
	s.player.SetPositionAtBottomLeft(image.Point{X: 0 + sideMargin, Y: s.height - bottomMargin})
}

func (s *ScenarioImpl) SetPlayerPositionAtRight() {
	s.player.SetPositionAtBottomRight(image.Point{X: s.width - sideMargin, Y: s.height - bottomMargin})
}
