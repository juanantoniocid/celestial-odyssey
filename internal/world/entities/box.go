package entities

import "image"

type Box struct {
	position      image.Point
	width, height int
}

func NewBox(position image.Point, width, height int) *Box {
	return &Box{
		position: position,
		width:    width,
		height:   height,
	}
}

func (b *Box) Bounds() image.Rectangle {
	return image.Rect(b.position.X, b.position.Y, b.position.X+b.width, b.position.Y+b.height)
}
