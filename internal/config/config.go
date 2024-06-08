package config

type Config struct {
	Window Window
	Screen Screen
	Player Player
	Ground Ground
}

func LoadConfig() Config {
	return Config{
		Window: setupWindow(),
		Screen: setupScreen(),
		Ground: setupGround(),
		Player: setupPlayer(),
	}
}
