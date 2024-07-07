package entity

type HorizontalDirection int

const (
	DirectionLeft HorizontalDirection = iota
	DirectionRight
)

type CharacterAction int

const (
	ActionIdle CharacterAction = iota
	ActionWalking
	ActionJumping
)
