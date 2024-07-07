package screen

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/entity"
)

const (
	sideMargin = 8
)

type ScenarioImpl struct {
	player *entity.Player
	world  *entity.World
	em     *entity.EntityManager

	renderer       Renderer
	inputHandler   InputHandler
	physicsHandler PhysicsHandler
}

func NewScenario(player *entity.Player, renderer Renderer, inputHandler InputHandler, physicsHandler PhysicsHandler, entityManager *entity.EntityManager, width int, height int) *ScenarioImpl {
	return &ScenarioImpl{
		player:         player,
		renderer:       renderer,
		inputHandler:   inputHandler,
		physicsHandler: physicsHandler,
		world:          entity.NewWorld(player, width, height),
		em:             entityManager,
	}
}

func (s *ScenarioImpl) Update() error {
	s.inputHandler.UpdatePlayer(s.player)
	s.player.Update()
	s.physicsHandler.ApplyPhysics(s.world, s.em.Entities())

	return nil
}

func (s *ScenarioImpl) Draw(screen *ebiten.Image) {
	s.renderer.Draw(screen, s.world, s.em.Entities())
}

func (s *ScenarioImpl) ShouldTransitionRight() bool {
	return s.player.Position().X+s.player.Width() >= s.world.GetWidth()
}

func (s *ScenarioImpl) ShouldTransitionLeft() bool {
	return s.player.Position().X <= 0
}

func (s *ScenarioImpl) SetPlayerPositionAtLeft() {
	s.player.SetPositionX(sideMargin)
}

func (s *ScenarioImpl) SetPlayerPositionAtRight() {
	s.player.SetPositionX(s.world.GetWidth() - sideMargin - s.player.Width())
}
