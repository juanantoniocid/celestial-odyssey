package config

import (
	"image/color"

	"celestial-odyssey/internal/util"
)

const (
	screenWidth  = 320
	screenHeight = 200
)

type Screen struct {
	ClearedEveryFrame bool
	Dimensions        util.Dimensions
	BackgroundColor   color.RGBA
}

func setupScreen() Screen {
	return Screen{
		ClearedEveryFrame: true,
		Dimensions:        util.Dimensions{Width: screenWidth, Height: screenHeight},
		BackgroundColor:   color.RGBA{R: 24, G: 8, B: 50, A: 1},
	}
}
