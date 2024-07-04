package entities

type World struct {
	width  int
	height int

	player  *Player
	grounds []*Ground
	boxes   []*Box
}

func NewWorld(player *Player, width, height int) *World {
	return &World{
		width:   width,
		height:  height,
		player:  player,
		grounds: []*Ground{},
		boxes:   []*Box{},
	}
}

func (w *World) AddGround(ground *Ground) {
	w.grounds = append(w.grounds, ground)
}

func (w *World) AddBox(box *Box) {
	w.boxes = append(w.boxes, box)
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

func (w *World) GetGrounds() []*Ground {
	return w.grounds
}

func (w *World) GetBoxes() []*Box {
	return w.boxes
}

func (w *World) GetCollidables() []Collidable {
	collidables := make([]Collidable, 0, len(w.grounds)+len(w.boxes))
	for _, g := range w.grounds {
		collidables = append(collidables, g)
	}
	for _, b := range w.boxes {
		collidables = append(collidables, b)
	}
	return collidables
}
