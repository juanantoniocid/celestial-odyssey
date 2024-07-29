package game

import (
	"celestial-odyssey/internal/legacy"
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/entity"
)

// Level represents a game level, which can contain multiple sections.
type Level interface {
	// Init initializes the level with the given parameters.
	Init()

	// AddSection adds a section to the level.
	AddSection(section Section)

	// CurrentSection returns the current section.
	CurrentSection() Section
}

// Section represents a single scenario within a level.
type Section interface {
	// Entities returns the entities in the scenario.
	Entities() *entity.Entities

	// SetPlayerPositionAtLeft sets the player position at the left side of the screen.
	SetPlayerPositionAtLeft()
}

// InputHandler is responsible for handling input.
type InputHandler interface {
	// UpdatePlayer updates the player based on the input.
	UpdatePlayer(player *legacy.Player)
}

// Renderer is responsible for drawing the game entities on the screen.
type Renderer interface {
	// Draw draws the game entities on the screen.
	Draw(screen *ebiten.Image, player *legacy.Player, entities *entity.Entities)
}

// CollisionHandler is responsible for applying physics to the game entities.
type CollisionHandler interface {
	// UpdatePlayer applies physics to the player and entities.
	UpdatePlayer(player *legacy.Player, collection *entity.Entities)
}
