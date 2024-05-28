package entities

import "image"

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

type Collidable interface {
	Bounds() image.Rectangle
}
