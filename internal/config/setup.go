package config

import (
	"celestial-odyssey/internal/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

const (
	screenWidth  = 320
	screenHeight = 200
)

func setupWindow() Window {
	return Window{
		Title:        "Celestial Odyssey",
		Dimensions:   util.Dimensions{Width: 960, Height: 600},
		ResizingMode: ebiten.WindowResizingModeEnabled,
	}
}

func setupScreen() Screen {
	return Screen{
		ClearedEveryFrame: true,
		Dimensions:        util.Dimensions{Width: screenWidth, Height: screenHeight},
		BackgroundColor:   color.RGBA{R: 24, G: 8, B: 50, A: 1},
	}
}

func setupGround() Ground {
	return Ground{
		Dimensions: util.Dimensions{Width: 40, Height: 28},
		File:       "assets/images/ground.png",
	}
}

func setupPlayer() Player {
	return Player{
		Dimensions:          util.Dimensions{Width: 16, Height: 32},
		WalkingVelocity:     2,
		InitialJumpVelocity: -10,
		Gravity:             0.5,
		SpritesFile:         "assets/images/player.png",
	}
}
