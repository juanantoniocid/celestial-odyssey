package screen

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/entity"
)

const (
	sideMargin = 8
)

type ScenarioImpl struct {
	player           *entity.Player
	character        *entity.GameEntity
	entityCollection *entity.Collection

	renderer         Renderer
	inputHandler     InputHandler
	collisionHandler CollisionHandler
	systemManager    SystemManager
}

func NewScenario(player *entity.Player, renderer Renderer, inputHandler InputHandler, collisionHandler CollisionHandler, character *entity.GameEntity, entityCollection *entity.Collection, systemManager SystemManager) *ScenarioImpl {
	return &ScenarioImpl{
		player:           player,
		character:        character,
		renderer:         renderer,
		inputHandler:     inputHandler,
		collisionHandler: collisionHandler,
		entityCollection: entityCollection,
		systemManager:    systemManager,
	}
}

func (s *ScenarioImpl) Update() error {
	s.systemManager.Update(s.entityCollection)
	s.inputHandler.UpdatePlayer(s.player)
	s.player.Update()
	s.collisionHandler.UpdatePlayer(s.player, s.entityCollection)

	return nil
}

func (s *ScenarioImpl) Draw(screen *ebiten.Image) {
	s.renderer.Draw(screen, s.player, s.character, s.entityCollection)
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
