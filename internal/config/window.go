package config

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Window struct {
	Title        string
	Dimensions   Dimensions
	ResizingMode ebiten.WindowResizingModeType
}

func setupWindow() Window {
	return Window{
		Title:        "Celestial Odyssey",
		Dimensions:   Dimensions{Width: 960, Height: 600},
		ResizingMode: ebiten.WindowResizingModeEnabled,
	}
}
