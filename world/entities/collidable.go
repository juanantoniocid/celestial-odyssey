package entities

import "image"

type Collidable interface {
	Bounds() image.Rectangle
}
