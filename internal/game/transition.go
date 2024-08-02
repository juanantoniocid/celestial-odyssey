package game

import "celestial-odyssey/internal/entity"

// TransitionCondition is a function that determines whether a transition should occur.
type TransitionCondition func(entities *entity.Entities) bool

// TransitionAction is a function that performs a transition.
type TransitionAction func(entities *entity.Entities)

// Transition represents a transition between two sections.
type Transition struct {
	fromSection SectionID
	toSection   SectionID
	condition   TransitionCondition
	action      TransitionAction
}

// NewTransition creates a new transition between two sections.
func NewTransition(from, to SectionID, condition TransitionCondition, action TransitionAction) *Transition {
	return &Transition{
		fromSection: from,
		toSection:   to,
		condition:   condition,
		action:      action,
	}
}
