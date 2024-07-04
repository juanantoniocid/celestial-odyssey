package physics

import (
	entities2 "celestial-odyssey/internal/entities"
	"log"
)

// PhysicsHandler is responsible for applying physics to the game entities.
type PhysicsHandler struct{}

// NewPhysicsHandler creates a new PhysicsHandler instance.
func NewPhysicsHandler() *PhysicsHandler {
	return &PhysicsHandler{}
}

// ApplyPhysics applies physics to the world entities.
func (h *PhysicsHandler) ApplyPhysics(world *entities2.World) {
	h.checkCollisions(world.GetPlayer(), world.GetCollidables())
	h.checkIfPlayerIsOnPlatform(world.GetPlayer(), world.GetCollidables(), world.GetHeight())
	h.enforceBoundaries(world.GetPlayer(), world.GetWidth(), world.GetHeight())
}

func (h *PhysicsHandler) checkCollisions(player *entities2.Player, collidables []entities2.Collidable) {
	for _, c := range collidables {
		if player.Bounds().Overlaps(c.Bounds()) {
			h.handleCollision(player, c)
		}
	}
}

func (h *PhysicsHandler) handleCollision(player *entities2.Player, c entities2.Collidable) {
	collision := player.Bounds().Intersect(c.Bounds())

	if collision.Dx() < collision.Dy() {
		handleHorizontalCollision(player, c)
		handleVerticalCollision(player, c)
	} else {
		handleVerticalCollision(player, c)
		handleHorizontalCollision(player, c)
	}
}

func handleVerticalCollision(player *entities2.Player, c entities2.Collidable) {
	if !player.Bounds().Overlaps(c.Bounds()) {
		return
	}

	playerBounds := player.Bounds()
	collidableBounds := c.Bounds()

	// Landing on top of the box
	if player.VerticalVelocity() > 0 &&
		playerBounds.Min.Y < collidableBounds.Min.Y &&
		playerBounds.Max.Y > collidableBounds.Min.Y {
		log.Println("Landing on top of the box")
		log.Println(playerBounds.Min.Y, collidableBounds.Min.Y, collidableBounds.Max.Y)
		player.SetPositionY(collidableBounds.Min.Y - player.Height())
		player.Land()
		return
	}

	// Hitting the bottom of the box
	if player.VerticalVelocity() < 0 {
		if playerBounds.Min.Y > collidableBounds.Min.Y &&
			playerBounds.Min.Y < collidableBounds.Max.Y {
			log.Println("Hitting the bottom of the box")
			player.SetPositionY(collidableBounds.Max.Y)
			player.Land()
		}
	}
}

func handleHorizontalCollision(player *entities2.Player, c entities2.Collidable) {
	if !player.Bounds().Overlaps(c.Bounds()) {
		return
	}

	playerBounds := player.Bounds()
	collidableBounds := c.Bounds()

	// Hitting the right side of the box
	if player.HorizontalVelocity() > 0 &&
		playerBounds.Min.X < collidableBounds.Min.X &&
		playerBounds.Max.X > collidableBounds.Min.X {
		player.SetPositionX(collidableBounds.Min.X - player.Width())
		player.Stop()
		log.Println("Hitting the right side of the box")
		return
	}

	// Hitting the left side of the box
	if player.HorizontalVelocity() < 0 &&
		playerBounds.Min.X < collidableBounds.Max.X &&
		playerBounds.Max.X > collidableBounds.Max.X {
		player.SetPositionX(collidableBounds.Max.X)
		player.Stop()
		log.Println("Hitting the left side of the box")
	}
}

func (h *PhysicsHandler) checkIfPlayerIsOnPlatform(player *entities2.Player, collidables []entities2.Collidable, height int) {
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

func (h *PhysicsHandler) enforceBoundaries(player *entities2.Player, width, height int) {
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
