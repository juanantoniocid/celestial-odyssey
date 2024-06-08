package config

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/util"
)

type Window struct {
	Title        string
	Dimensions   util.Dimensions
	ResizingMode ebiten.WindowResizingModeType
}

func setupWindow() Window {
	return Window{
		Title:        "Celestial Odyssey",
		Dimensions:   util.Dimensions{Width: 960, Height: 600},
		ResizingMode: ebiten.WindowResizingModeEnabled,
	}
}
