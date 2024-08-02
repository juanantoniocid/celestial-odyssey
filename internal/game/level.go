package game

import (
	"celestial-odyssey/internal/entity"
)

// Section represents a single scenario within a level.
type Section interface {
	// Entities returns the entities in the scenario.
	Entities() *entity.Entities
}

// BasicLevel represents a basic level in the game.
type BasicLevel struct {
	currentSectionID SectionID
	sections         map[SectionID]Section
	transitions      []*Transition
}

// NewBasicLevel creates a new basic level.
func NewBasicLevel() *BasicLevel {
	return &BasicLevel{
		sections:    make(map[SectionID]Section),
		transitions: make([]*Transition, 0),
	}
}

// Update updates the level.
func (l *BasicLevel) Update() {
	currentSection := l.CurrentSection()
	entities := currentSection.Entities()

	for _, transition := range l.transitions {
		if transition.fromSection == l.currentSectionID {
			if transition.condition(entities) {
				transition.action(entities)
				l.SetCurrentSection(transition.toSection)

				break
			}
		}
	}
}

// Entities returns the entities in the level.
func (l *BasicLevel) Entities() *entity.Entities {
	return l.CurrentSection().Entities()
}

// AddSection adds a section to the level.
func (l *BasicLevel) AddSection(id SectionID, section Section) {
	l.sections[id] = section
}

// AddTransition adds a transition to the level.
func (l *BasicLevel) AddTransition(t *Transition) {
	l.transitions = append(l.transitions, t)
}

// CurrentSection returns the current section.
func (l *BasicLevel) CurrentSection() Section {
	section, found := l.sections[l.currentSectionID]
	if !found {
		panic("current section not found")
	}

	return section
}

// SetCurrentSection sets the current section.
func (l *BasicLevel) SetCurrentSection(id SectionID) {
	l.currentSectionID = id
}
