package entity

type World struct {
	width  int
	height int

	player *Player
}

func NewWorld(player *Player, width, height int) *World {
	return &World{
		width:  width,
		height: height,
		player: player,
	}
}

func (w *World) GetWidth() int {
	return w.width
}

func (w *World) GetHeight() int {
	return w.height
}

func (w *World) GetPlayer() *Player {
	return w.player
}
