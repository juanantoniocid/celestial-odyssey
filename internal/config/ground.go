package config

type Ground struct {
	File string
}

func setupGround() Ground {
	return Ground{
		File: "internal/assets/ground.png",
	}
}
