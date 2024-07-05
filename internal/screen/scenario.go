package screen

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/entities"
)

const (
	sideMargin = 8
)

type ScenarioImpl struct {
	player *entities.Player
	world  *entities.World
	em     *entities.EntityManager

	renderer       Renderer
	inputHandler   InputHandler
	physicsHandler PhysicsHandler
}

func NewScenario(player *entities.Player, renderer Renderer, inputHandler InputHandler, physicsHandler PhysicsHandler, width int, height int) *ScenarioImpl {
	return &ScenarioImpl{
		player:         player,
		renderer:       renderer,
		inputHandler:   inputHandler,
		physicsHandler: physicsHandler,
		world:          entities.NewWorld(player, width, height),
		em:             entities.NewEntityManager(),
	}
}

func (s *ScenarioImpl) CreateEntity() *entities.Entity {
	return s.em.CreateEntity()
}

func (s *ScenarioImpl) AddBox(b *entities.Box) {
	s.world.AddBox(b)
}

func (s *ScenarioImpl) AddGround(g *entities.Ground) {
	s.world.AddGround(g)
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
