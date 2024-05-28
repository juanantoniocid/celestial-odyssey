package entities

import "image"

type Platform struct {
	position      image.Point
	width, height int
}

func NewPlatform(position image.Point, width, height int) *Platform {
	return &Platform{
		position: position,
		width:    width,
		height:   height,
	}
}

func (p *Platform) Bounds() image.Rectangle {
	return image.Rect(p.position.X, p.position.Y, p.position.X+p.width, p.position.Y+p.height)
}
