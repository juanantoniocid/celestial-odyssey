package config

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/util"
)

const (
	screenWidth  = 320
	screenHeight = 200
)

type Config struct {
	Window Window
	Screen Screen
	Player Player
}

type Window struct {
	Title        string
	Dimensions   util.Dimensions
	ResizingMode ebiten.WindowResizingModeType
}

type Screen struct {
	ClearedEveryFrame bool
	Dimensions        util.Dimensions
}

type Player struct {
	Dimensions          util.Dimensions
	Speed               int
	InitialJumpVelocity float64
	Gravity             float64
	PlayArea            image.Rectangle
	SpritesFile         string
}

func LoadConfig() Config {
	return Config{
		Window: loadWindow(),
		Screen: loadScreen(),
		Player: loadPlayer(),
	}
}

func loadWindow() Window {
	return Window{
		Title:        "Celestial Odyssey",
		Dimensions:   util.Dimensions{Width: 960, Height: 600},
		ResizingMode: ebiten.WindowResizingModeEnabled,
	}
}

func loadScreen() Screen {
	return Screen{
		ClearedEveryFrame: true,
		Dimensions:        util.Dimensions{Width: screenWidth, Height: screenHeight},
	}
}

func loadPlayer() Player {
	return Player{
		Dimensions:          util.Dimensions{Width: 16, Height: 32},
		Speed:               2,
		InitialJumpVelocity: -10,
		Gravity:             0.5,
		PlayArea: image.Rectangle{
			Min: image.Point{X: 0, Y: 0},
			Max: image.Point{X: screenWidth - 1, Y: screenHeight - 1},
		},
		SpritesFile: "assets/images/player.png",
	}
}
