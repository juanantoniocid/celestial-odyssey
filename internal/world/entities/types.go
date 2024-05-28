package entities

import "image"

type HorizontalDirection int

const (
	Left HorizontalDirection = iota
	Right
)

type Action int

const (
	Idle Action = iota
	Walking
	Jumping
)

type Collidable interface {
	Bounds() image.Rectangle
}
