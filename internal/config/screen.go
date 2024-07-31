package config

import (
	"image/color"
)

const (
	ScreenWidth  = 320
	ScreenHeight = 200
)

type Screen struct {
	ClearedEveryFrame bool
	Dimensions        Dimensions
	BackgroundColor   color.RGBA
}

func setupScreen() Screen {
	return Screen{
		ClearedEveryFrame: true,
		Dimensions:        Dimensions{Width: ScreenWidth, Height: ScreenHeight},
		BackgroundColor:   color.RGBA{R: 24, G: 8, B: 50, A: 1},
	}
}
