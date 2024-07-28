package game

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/system/behavior"
	"celestial-odyssey/internal/system/graphics"
)

type LevelImpl struct {
	sections       []Section
	currentSection int

	updateSystem behavior.UpdateSystem
	renderer     graphics.Renderer
}

func NewLevel(updateSystem behavior.UpdateSystem, renderer graphics.Renderer) *LevelImpl {
	return &LevelImpl{
		sections:       make([]Section, 0),
		currentSection: 0,
		updateSystem:   updateSystem,
		renderer:       renderer,
	}
}

func (l *LevelImpl) AddSection(section Section) {
	l.sections = append(l.sections, section)
}

func (l *LevelImpl) Init() {
	currentSection := l.sections[l.currentSection]
	currentSection.SetPlayerPositionAtLeft()
}

func (l *LevelImpl) Update() error {
	currentSection := l.sections[l.currentSection]
	entities := currentSection.Entities()
	l.updateSystem.Update(entities)

	if currentSection.ShouldTransitionRight() && l.currentSection < len(l.sections)-1 {
		l.currentSection++
		nextScenario := l.sections[l.currentSection]
		nextScenario.SetPlayerPositionAtLeft()
	} else if currentSection.ShouldTransitionLeft() && l.currentSection > 0 {
		l.currentSection--
		prevScenario := l.sections[l.currentSection]
		prevScenario.SetPlayerPositionAtRight()
	}

	return nil
}

func (l *LevelImpl) Draw(screen *ebiten.Image) {
	currentSection := l.sections[l.currentSection]
	entities := currentSection.Entities()
	l.renderer.Draw(screen, entities)
}
