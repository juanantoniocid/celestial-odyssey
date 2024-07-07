package systems

import (
	"log"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/entity"
)

// CollisionHandler is responsible for applying physics to the game entities.
type CollisionHandler struct{}

// NewCollisionHandler creates a new CollisionHandler instance.
func NewCollisionHandler() *CollisionHandler {
	return &CollisionHandler{}
}

// Update updates the game entities based on the collision rules.
func (h *CollisionHandler) Update(player *entity.Player, entities []*entity.GameEntity) {
	collidables := make([]entity.Collidable, 0)
	for _, e := range entities {
		collidable := e.Bounds()
		collidables = append(collidables, collidable)
	}

	h.checkCollisions(player, collidables)
	h.checkIfPlayerIsOnPlatform(player, collidables, config.ScreenHeight)
	h.enforceBoundaries(player, config.ScreenWidth, config.ScreenHeight)
}

func (h *CollisionHandler) checkCollisions(player *entity.Player, collidables []entity.Collidable) {
	for _, c := range collidables {
		if player.Bounds().Overlaps(c.Bounds()) {
			h.handleCollision(player, c)
		}
	}
}

func (h *CollisionHandler) handleCollision(player *entity.Player, c entity.Collidable) {
	collision := player.Bounds().Intersect(c.Bounds())

	if collision.Dx() < collision.Dy() {
		handleHorizontalCollision(player, c)
		handleVerticalCollision(player, c)
	} else {
		handleVerticalCollision(player, c)
		handleHorizontalCollision(player, c)
	}
}

func handleVerticalCollision(player *entity.Player, c entity.Collidable) {
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

func handleHorizontalCollision(player *entity.Player, c entity.Collidable) {
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

func (h *CollisionHandler) checkIfPlayerIsOnPlatform(player *entity.Player, collidables []entity.Collidable, height int) {
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

func (h *CollisionHandler) enforceBoundaries(player *entity.Player, width, height int) {
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
