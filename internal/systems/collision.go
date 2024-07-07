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
	h.checkCollisions(player, entities)
	h.checkIfPlayerIsOnPlatform(player, entities, config.ScreenHeight)
	h.enforceBoundaries(player, config.ScreenWidth, config.ScreenHeight)
}

func (h *CollisionHandler) checkCollisions(player *entity.Player, entities []*entity.GameEntity) {
	for _, e := range entities {
		if player.Bounds().Overlaps(e.Bounds()) {
			h.handleCollision(player, e)
		}
	}
}

func (h *CollisionHandler) handleCollision(player *entity.Player, entity *entity.GameEntity) {
	collision := player.Bounds().Intersect(entity.Bounds())

	if collision.Dx() < collision.Dy() {
		handleHorizontalCollision(player, entity)
		handleVerticalCollision(player, entity)
	} else {
		handleVerticalCollision(player, entity)
		handleHorizontalCollision(player, entity)
	}
}

func handleVerticalCollision(player *entity.Player, entity *entity.GameEntity) {
	if !player.Bounds().Overlaps(entity.Bounds()) {
		return
	}

	playerBounds := player.Bounds()
	collidableBounds := entity.Bounds()

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

func handleHorizontalCollision(player *entity.Player, entity *entity.GameEntity) {
	if !player.Bounds().Overlaps(entity.Bounds()) {
		return
	}

	playerBounds := player.Bounds()
	collidableBounds := entity.Bounds()

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

func (h *CollisionHandler) checkIfPlayerIsOnPlatform(player *entity.Player, entities []*entity.GameEntity, height int) {
	playerBounds := player.Bounds()
	isOnPlatform := false

	// Check if the player is on any platform
	for _, e := range entities {
		entityBounds := e.Bounds()

		// Check if the player is on top of the box
		if playerBounds.Max.Y == entityBounds.Min.Y &&
			playerBounds.Min.X < entityBounds.Max.X &&
			playerBounds.Max.X > entityBounds.Min.X {
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
