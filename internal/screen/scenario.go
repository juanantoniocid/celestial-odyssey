package screen

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/entity"
	"celestial-odyssey/internal/system"
	"celestial-odyssey/internal/system/graphics"
)

const (
	sideMargin = 8
)

type ScenarioImpl struct {
	entities *entity.Entities
	systems  system.System
	renderer graphics.Renderer
}

func NewScenario(entities *entity.Entities, systems system.System, renderer graphics.Renderer) *ScenarioImpl {
	return &ScenarioImpl{
		entities: entities,
		systems:  systems,
		renderer: renderer,
	}
}

func (s *ScenarioImpl) Update() error {
	s.systems.Update(s.entities)

	return nil
}

func (s *ScenarioImpl) Draw(screen *ebiten.Image) {
	s.renderer.Draw(screen, s.entities)
}

func (s *ScenarioImpl) ShouldTransitionRight() bool {
	//return s.player.Position().X+s.player.Width() >= config.ScreenWidth
	return false
}

func (s *ScenarioImpl) ShouldTransitionLeft() bool {
	// return s.player.Position().X <= 0
	return false
}

func (s *ScenarioImpl) SetPlayerPositionAtLeft() {
	// s.player.SetPositionX(sideMargin)
}

func (s *ScenarioImpl) SetPlayerPositionAtRight() {
	// s.player.SetPositionX(config.ScreenWidth - sideMargin - s.player.Width())
}
