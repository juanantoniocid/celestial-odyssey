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
	player    *entity.Player
	character *entity.GameEntity
	entities  *entity.Collection

	renderer           Renderer
	inputHandler       InputHandler
	collisionHandler   CollisionHandler
	systemInputHandler SystemInputHandler
}

func NewScenario(player *entity.Player, renderer Renderer, inputHandler InputHandler, collisionHandler CollisionHandler, systemInputHandler SystemInputHandler, character *entity.GameEntity, entities *entity.Collection) *ScenarioImpl {
	return &ScenarioImpl{
		player:             player,
		character:          character,
		renderer:           renderer,
		inputHandler:       inputHandler,
		collisionHandler:   collisionHandler,
		systemInputHandler: systemInputHandler,
		entities:           entities,
	}
}

func (s *ScenarioImpl) Update() error {
	s.inputHandler.UpdateCharacter(s.character)
	s.systemInputHandler.Update(s.character)
	s.inputHandler.UpdatePlayer(s.player)
	s.player.Update()
	s.collisionHandler.Update(s.player, s.entities.GameEntities())

	return nil
}

func (s *ScenarioImpl) Draw(screen *ebiten.Image) {
	s.renderer.Draw(screen, s.player, s.character, s.entities.GameEntities())
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
