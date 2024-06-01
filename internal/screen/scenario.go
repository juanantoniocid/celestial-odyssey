package screen

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/world/entities"
)

const (
	sideMargin = 8
)

type InputHandler interface {
	UpdatePlayer(player *entities.Player)
}

type Renderer interface {
	DrawBackground(screen *ebiten.Image, background *ebiten.Image)
	DrawPlayer(screen *ebiten.Image, player *entities.Player)
	DrawCollidable(screen *ebiten.Image, collidable entities.Collidable)
}

type PhysicsHandler interface {
	Update(player *entities.Player, collidables []entities.Collidable, width, height int)
}

type ScenarioImpl struct {
	player      *entities.Player
	collidables []entities.Collidable

	background     *ebiten.Image
	renderer       Renderer
	inputHandler   InputHandler
	physicsHandler PhysicsHandler

	width  int
	height int
}

func NewScenario(player *entities.Player, background *ebiten.Image, renderer Renderer, inputHandler InputHandler, physicsHandler PhysicsHandler, width int, height int) *ScenarioImpl {
	return &ScenarioImpl{
		player:         player,
		background:     background,
		renderer:       renderer,
		inputHandler:   inputHandler,
		physicsHandler: physicsHandler,
		width:          width,
		height:         height,
		collidables:    make([]entities.Collidable, 0),
	}
}

func (s *ScenarioImpl) AddCollidable(c entities.Collidable) {
	s.collidables = append(s.collidables, c)
}

func (s *ScenarioImpl) Update() error {
	s.inputHandler.UpdatePlayer(s.player)
	s.player.Update()

	s.physicsHandler.Update(s.player, s.collidables, s.width, s.height)

	return nil
}

func (s *ScenarioImpl) Draw(screen *ebiten.Image) {
	s.renderer.DrawBackground(screen, s.background)
	s.renderer.DrawPlayer(screen, s.player)
	for _, c := range s.collidables {
		s.renderer.DrawCollidable(screen, c)
	}
}

func (s *ScenarioImpl) ShouldTransitionRight() bool {
	return s.player.Position().X+s.player.Width() >= s.width
}

func (s *ScenarioImpl) ShouldTransitionLeft() bool {
	return s.player.Position().X <= 0
}

func (s *ScenarioImpl) SetPlayerPositionAtLeft() {
	s.player.SetPositionX(sideMargin)
}

func (s *ScenarioImpl) SetPlayerPositionAtRight() {
	s.player.SetPositionX(s.width - sideMargin - s.player.Width())
}
