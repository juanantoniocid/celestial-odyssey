package screen

import (
	"celestial-odyssey/internal/system"
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/entity"
)

const (
	sideMargin = 8
)

type ScenarioImpl struct {
	player   *entity.Player
	renderer Renderer

	inputHandler     InputHandler
	collisionHandler CollisionHandler

	entities *entity.Entities
	systems  system.System
}

func NewScenario(player *entity.Player, renderer Renderer, inputHandler InputHandler, collisionHandler CollisionHandler, entities *entity.Entities, systems system.System) *ScenarioImpl {
	return &ScenarioImpl{
		player:           player,
		entities:         entities,
		renderer:         renderer,
		inputHandler:     inputHandler,
		collisionHandler: collisionHandler,
		systems:          systems,
	}
}

func (s *ScenarioImpl) Update() error {
	s.systems.Update(s.entities)
	s.inputHandler.UpdatePlayer(s.player)
	s.player.Update()
	s.collisionHandler.UpdatePlayer(s.player, s.entities)

	return nil
}

func (s *ScenarioImpl) Draw(screen *ebiten.Image) {
	s.renderer.Draw(screen, s.player, s.entities)
}

func (s *ScenarioImpl) ShouldTransitionRight() bool {
	return s.player.Position().X+s.player.Width() >= config.ScreenWidth
}

func (s *ScenarioImpl) ShouldTransitionLeft() bool {
	return s.player.Position().X <= 0
}

func (s *ScenarioImpl) SetPlayerPositionAtLeft() {
	s.player.SetPositionX(sideMargin)
}

func (s *ScenarioImpl) SetPlayerPositionAtRight() {
	s.player.SetPositionX(config.ScreenWidth - sideMargin - s.player.Width())
}
