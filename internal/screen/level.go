package screen

import "github.com/hajimehoshi/ebiten/v2"

type Level interface {
	Update()
	Draw(screen *ebiten.Image)
}
