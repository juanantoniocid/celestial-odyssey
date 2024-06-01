package screen

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/world/entities"
)

type InputHandler interface {
	UpdatePlayer(player *entities.Player)
}

type Renderer interface {
	DrawBackground(screen *ebiten.Image, background *ebiten.Image)
	DrawPlayer(screen *ebiten.Image, player *entities.Player)
	DrawCollidable(screen *ebiten.Image, collidable entities.Collidable)
}

type PhysicsHandler interface {
	Update(player *entities.Player, collidables []entities.Collidable, width, height int)
}
