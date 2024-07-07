package config

import (
	"image/color"

	"celestial-odyssey/internal/util"
)

const (
	ScreenWidth  = 320
	ScreenHeight = 200
)

type Screen struct {
	ClearedEveryFrame bool
	Dimensions        util.Dimensions
	BackgroundColor   color.RGBA
}

func setupScreen() Screen {
	return Screen{
		ClearedEveryFrame: true,
		Dimensions:        util.Dimensions{Width: ScreenWidth, Height: ScreenHeight},
		BackgroundColor:   color.RGBA{R: 24, G: 8, B: 50, A: 1},
	}
}
