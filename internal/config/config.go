package config

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/util"
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
		Window: setupWindow(),
		Screen: setupScreen(),
		Ground: setupGround(),
		Player: setupPlayer(),
	}
}
