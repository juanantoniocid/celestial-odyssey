package config

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/util"
)

const (
	screenWidth  = 320
	screenHeight = 200
)

type Config struct {
	Window Window
	Screen Screen
	Player Player
	Ground Ground
}

type Window struct {
	Title        string
	Dimensions   util.Dimensions
	ResizingMode ebiten.WindowResizingModeType
}

type Screen struct {
	ClearedEveryFrame bool
	Dimensions        util.Dimensions
	BackgroundColor   color.RGBA
}

type Ground struct {
	Dimensions util.Dimensions
	File       string
}

type Player struct {
	Dimensions          util.Dimensions
	WalkingVelocity     int
	InitialJumpVelocity float64
	Gravity             float64
	SpritesFile         string
}

func LoadConfig() Config {
	return Config{
		Window: loadWindow(),
		Screen: loadScreen(),
		Ground: loadGround(),
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
		BackgroundColor:   color.RGBA{R: 24, G: 8, B: 50, A: 1},
	}
}

func loadGround() Ground {
	return Ground{
		Dimensions: util.Dimensions{Width: 40, Height: 28},
		File:       "assets/images/ground.png",
	}
}

func loadPlayer() Player {
	return Player{
		Dimensions:          util.Dimensions{Width: 16, Height: 32},
		WalkingVelocity:     2,
		InitialJumpVelocity: -10,
		Gravity:             0.5,
		SpritesFile:         "assets/images/player.png",
	}
}
