package entities

import "image"

type Box struct {
	image.Rectangle
}

func NewBox(rectangle image.Rectangle) *Box {
	return &Box{
		Rectangle: rectangle,
	}
}
