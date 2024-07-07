package screen

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/entity"
)

// Level represents a game level, which can contain multiple scenarios.
// Each level is responsible for initializing itself, managing its scenarios,
// updating its state, and drawing itself on the screen.
type Level interface {
	// Init initializes the level with the given parameters.
	Init()

	// AddScenario adds a scenario to the level.
	// This allows the level to manage multiple scenarios within it.
	AddScenario(scenario Scenario)

	// Update updates the level state.
	Update() error

	// Draw draws the level on the screen.
	Draw(screen *ebiten.Image)
}

// Scenario represents a single scenario within a level.
// Each scenario is responsible for updating its state, drawing itself on the screen,
// and handling transitions within the level.
type Scenario interface {
	// Update updates the scenario state.
	Update() error

	// Draw draws the scenario on the screen.
	Draw(screen *ebiten.Image)

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
	UpdatePlayer(player *entity.Player)
}

// Renderer is responsible for drawing the game entities on the screen.
type Renderer interface {
	// Draw draws the game entities on the screen.
	Draw(screen *ebiten.Image, player *entity.Player, entities map[entity.ID]*entity.GameEntity)
}

// PhysicsHandler is responsible for applying physics to the game entities.
type PhysicsHandler interface {
	// ApplyPhysics applies physics to the player and collidables.
	ApplyPhysics(player *entity.Player, entities map[entity.ID]*entity.GameEntity)
}
