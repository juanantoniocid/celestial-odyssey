package entity

import "image"

type Ground struct {
	image.Rectangle
}

func NewGround(rectangle image.Rectangle) *Ground {
	return &Ground{
		Rectangle: rectangle,
	}
}
