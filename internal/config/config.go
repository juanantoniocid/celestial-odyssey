package config

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/util"
)

type Config struct {
	Window Window
	Screen Screen
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

func LoadConfig() Config {
	return Config{
		Window: loadWindow(),
		Screen: loadScreen(),
	}
}

func loadWindow() Window {
	return Window{
		Title:        "Celestial Odyssey",
		Dimensions:   util.Dimensions{Width: 960, Height: 720},
		ResizingMode: ebiten.WindowResizingModeEnabled,
	}
}

func loadScreen() Screen {
	return Screen{
		ClearedEveryFrame: true,
		Dimensions:        util.Dimensions{Width: 320, Height: 240},
	}
}
