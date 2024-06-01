package physics

import "celestial-odyssey/internal/world/entities"

// PhysicsHandler is responsible for applying physics to the game entities.
type PhysicsHandler struct{}

// NewPhysicsHandler creates a new PhysicsHandler instance.
func NewPhysicsHandler() *PhysicsHandler {
	return &PhysicsHandler{}
}

// ApplyPhysics applies physics to the player and collidables.
func (h *PhysicsHandler) ApplyPhysics(player *entities.Player, collidables []entities.Collidable, width, height int) {
	h.checkCollisions(player, collidables)
	h.checkIfPlayerIsOnPlatform(player, collidables, height)
	h.enforceBoundaries(player, width, height)
}

func (h *PhysicsHandler) checkCollisions(player *entities.Player, collidables []entities.Collidable) {
	for _, c := range collidables {
		if player.Bounds().Overlaps(c.Bounds()) {
			h.handleCollision(player, c)
		}
	}
}

func (h *PhysicsHandler) handleCollision(player *entities.Player, c entities.Collidable) {
	playerBounds := player.Bounds()
	collidableBounds := c.Bounds()

	// Check for vertical collisions first (jumping and landing)
	if playerBounds.Min.Y < collidableBounds.Max.Y && playerBounds.Max.Y > collidableBounds.Min.Y {
		// Landing on top of the box
		if player.Position().Y < collidableBounds.Min.Y && playerBounds.Max.Y > collidableBounds.Min.Y {
			player.SetPositionY(collidableBounds.Min.Y - player.Height())
			player.Land()
			return
		}
		// Hitting the bottom of the box
		if player.Position().Y > collidableBounds.Max.Y && playerBounds.Min.Y < collidableBounds.Max.Y {
			player.SetPositionY(collidableBounds.Max.Y)
			player.Land()
			return
		}
	}

	// Horizontal collision
	if playerBounds.Min.X < collidableBounds.Max.X && playerBounds.Max.X > collidableBounds.Min.X {
		// Hitting the right side of the box
		if playerBounds.Min.X < collidableBounds.Min.X && playerBounds.Max.X > collidableBounds.Min.X {
			player.SetPositionX(collidableBounds.Min.X - player.Width())
			player.Stop()
			return
		}

		// Hitting the left side of the box
		if playerBounds.Min.X < collidableBounds.Max.X && playerBounds.Max.X > collidableBounds.Max.X {
			player.SetPositionX(collidableBounds.Max.X)
			player.Stop()
			return
		}
	}
}

func (h *PhysicsHandler) checkIfPlayerIsOnPlatform(player *entities.Player, collidables []entities.Collidable, height int) {
	playerBounds := player.Bounds()
	isOnPlatform := false

	// Check if the player is on any platform
	for _, c := range collidables {
		collidableBounds := c.Bounds()

		// Check if the player is on top of the box
		if playerBounds.Max.Y == collidableBounds.Min.Y &&
			playerBounds.Min.X < collidableBounds.Max.X &&
			playerBounds.Max.X > collidableBounds.Min.X {
			isOnPlatform = true
			break
		}
	}

	// Check if the player is on the ground
	if playerBounds.Max.Y >= height {
		isOnPlatform = true
	}

	// If the player is not on any platform and not jumping, apply gravity
	if !isOnPlatform && !player.IsJumping() {
		player.Fall()
	}
}

func (h *PhysicsHandler) enforceBoundaries(player *entities.Player, width, height int) {
	if player.Position().X < 0 {
		player.SetPositionX(0)
	} else if player.Position().X+player.Width() > width {
		player.SetPositionX(width - player.Width())
	}

	if player.Position().Y < 0 {
		player.SetPositionY(0)
	} else if player.Position().Y+player.Height() > height {
		player.SetPositionY(height - player.Height())
		player.Land()
	}
}
