package config

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/util"
)

type Config struct {
	Window Window
	Game   Game
}

type Window struct {
	Title                   string
	Dimensions              util.Dimensions
	ResizingMode            ebiten.WindowResizingModeType
	ScreenClearedEveryFrame bool
}

type Game struct {
	Dimensions util.Dimensions
}

func LoadConfig() Config {
	return Config{
		Window: loadWindow(),
		Game:   loadGame(),
	}
}

func loadWindow() Window {
	return Window{
		Title:                   "Celestial Odyssey",
		Dimensions:              util.Dimensions{Width: 960, Height: 720},
		ResizingMode:            ebiten.WindowResizingModeEnabled,
		ScreenClearedEveryFrame: true,
	}
}

func loadGame() Game {
	return Game{
		Dimensions: util.Dimensions{Width: 320, Height: 240},
	}
}
