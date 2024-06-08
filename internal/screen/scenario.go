package screen

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/world/entities"
)

const (
	sideMargin = 8
)

type ScenarioImpl struct {
	player      *entities.Player
	ground      []*entities.Ground
	collidables []entities.Collidable

	renderer       Renderer
	inputHandler   InputHandler
	physicsHandler PhysicsHandler

	width  int
	height int
}

func NewScenario(player *entities.Player, renderer Renderer, inputHandler InputHandler, physicsHandler PhysicsHandler, width int, height int) *ScenarioImpl {
	return &ScenarioImpl{
		player:         player,
		renderer:       renderer,
		inputHandler:   inputHandler,
		physicsHandler: physicsHandler,
		width:          width,
		height:         height,

		ground:      make([]*entities.Ground, 0),
		collidables: make([]entities.Collidable, 0),
	}
}

func (s *ScenarioImpl) AddGround(g *entities.Ground) {
	s.ground = append(s.ground, g)
}

func (s *ScenarioImpl) AddCollidable(c entities.Collidable) {
	s.collidables = append(s.collidables, c)
}

func (s *ScenarioImpl) Update() error {
	s.inputHandler.UpdatePlayer(s.player)
	s.player.Update()

	s.physicsHandler.ApplyPhysics(s.player, s.collidables, s.width, s.height)

	return nil
}

func (s *ScenarioImpl) Draw(screen *ebiten.Image) {
	s.renderer.DrawBackground(screen, s.width, s.height)
	s.renderer.DrawPlayer(screen, s.player)
	s.renderer.DrawGround(screen, s.ground)
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
