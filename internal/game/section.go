package game

import (
	"celestial-odyssey/internal/entity"
)

const (
	sideMargin = 8
)

type BasicSection struct {
	entities *entity.Entities
}

func NewBasicSection(entities *entity.Entities) *BasicSection {
	return &BasicSection{
		entities: entities,
	}
}

func (s *BasicSection) Entities() *entity.Entities {
	return s.entities
}

func (s *BasicSection) ShouldTransitionRight() bool {
	//return s.player.Position().X+s.player.Width() >= config.ScreenWidth
	return false
}

func (s *BasicSection) ShouldTransitionLeft() bool {
	// return s.player.Position().X <= 0
	return false
}

func (s *BasicSection) SetPlayerPositionAtLeft() {
	// s.player.SetPositionX(sideMargin)
}

func (s *BasicSection) SetPlayerPositionAtRight() {
	// s.player.SetPositionX(config.ScreenWidth - sideMargin - s.player.Width())
}
