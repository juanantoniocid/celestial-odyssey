package config

import "celestial-odyssey/internal/util"

type Ground struct {
	Dimensions util.Dimensions
	File       string
}

func setupGround() Ground {
	return Ground{
		Dimensions: util.Dimensions{Width: 40, Height: 28},
		File:       "assets/images/ground.png",
	}
}
