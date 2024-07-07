package config

import "celestial-odyssey/internal/util"

type Player struct {
	Dimensions          util.Dimensions
	WalkingVelocity     int
	InitialJumpVelocity float64
	Gravity             float64
	SpriteSheet         string
}

func setupPlayer() Player {
	return Player{
		Dimensions:          util.Dimensions{Width: 16, Height: 32},
		WalkingVelocity:     2,
		InitialJumpVelocity: -10,
		Gravity:             0.5,
		SpriteSheet:         "internal/assets/player.png",
	}
}
