package entity

type World struct {
	player *Player
}

func NewWorld(player *Player) *World {
	return &World{
		player: player,
	}
}

func (w *World) GetPlayer() *Player {
	return w.player
}
