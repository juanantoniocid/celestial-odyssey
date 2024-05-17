package config

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"

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
	Speed       int
	PlayArea    image.Rectangle
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
		Dimensions:   util.Dimensions{Width: 960, Height: 600},
		ResizingMode: ebiten.WindowResizingModeEnabled,
	}
}

func loadScreen() Screen {
	return Screen{
		ClearedEveryFrame: true,
		Dimensions:        util.Dimensions{Width: 320, Height: 200},
	}
}

func loadPlayer() Player {
	return Player{
		Dimensions:  util.Dimensions{Width: 16, Height: 32},
		Speed:       2,
		PlayArea:    image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: 320 - 1, Y: 200 - 1}},
		SpritesFile: "assets/images/player.png",
	}
}
