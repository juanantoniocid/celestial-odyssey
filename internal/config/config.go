package config

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/util"
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
	Dimensions  util.Dimensions
	SpritesFile string
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

func loadPlayer() Player {
	return Player{
		Dimensions:  util.Dimensions{Width: 16, Height: 32},
		SpritesFile: "assets/images/player.png",
	}
}
