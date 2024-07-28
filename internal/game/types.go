package game

import (
	"celestial-odyssey/internal/legacy"
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/entity"
)

// Level represents a game level, which can contain multiple sections.
// Each level is responsible for initializing itself, managing its sections,
// updating its state, and drawing itself on the screen.
type Level interface {
	// Init initializes the level with the given parameters.
	Init()

	// AddSection adds a section to the level.
	AddSection(section Section)

	// Update updates the level state.
	Update() error

	// Draw draws the level on the screen.
	Draw(screen *ebiten.Image)
}

// Section represents a single scenario within a level.
// Each scenario is responsible for updating its state, drawing itself on the screen,
// and handling transitions within the level.
type Section interface {
	// Entities returns the entities in the scenario.
	Entities() *entity.Entities

	// ShouldTransitionRight returns true if the scenario should transition to the next scenario on the right.
	ShouldTransitionRight() bool

	// ShouldTransitionLeft returns true if the scenario should transition to the previous scenario on the left.
	ShouldTransitionLeft() bool

	// SetPlayerPositionAtLeft sets the player position on the left side of the scenario.
	SetPlayerPositionAtLeft()

	// SetPlayerPositionAtRight sets the player position on the right side of the scenario.
	SetPlayerPositionAtRight()
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
