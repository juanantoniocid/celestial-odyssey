package screen

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type LevelImpl struct {
	scenarios       []Scenario
	currentScenario int
}

func NewLevel() *LevelImpl {
	return &LevelImpl{
		scenarios:       make([]Scenario, 0),
		currentScenario: 0,
	}
}

func (l *LevelImpl) AddScenario(scenario Scenario) {
	l.scenarios = append(l.scenarios, scenario)
}

func (l *LevelImpl) Init() {
	currentScenario := l.scenarios[l.currentScenario]
	currentScenario.SetPlayerPositionAtLeft()
}

func (l *LevelImpl) Update() error {
	currentScenario := l.scenarios[l.currentScenario]
	err := currentScenario.Update()
	if err != nil {
		return err
	}

	if currentScenario.ShouldTransitionRight() && l.currentScenario < len(l.scenarios)-1 {
		l.currentScenario++
		nextScenario := l.scenarios[l.currentScenario]
		nextScenario.SetPlayerPositionAtLeft()
	} else if currentScenario.ShouldTransitionLeft() && l.currentScenario > 0 {
		l.currentScenario--
		prevScenario := l.scenarios[l.currentScenario]
		prevScenario.SetPlayerPositionAtRight()
	}

	return nil
}

func (l *LevelImpl) Draw(screen *ebiten.Image) {
	currentScenario := l.scenarios[l.currentScenario]
	currentScenario.Draw(screen)
}
