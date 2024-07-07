package systems

import (
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
	h.handleCollisions(player, entities)
	h.checkIfPlayerIsOnPlatform(player, entities, config.ScreenHeight)
	h.enforceBoundaries(player, config.ScreenWidth, config.ScreenHeight)
}

func (h *CollisionHandler) handleCollisions(player *entity.Player, entities []*entity.GameEntity) {
	for _, e := range entities {
		if player.Bounds().Overlaps(e.Bounds()) {
			h.handleCollision(player, e)
		}
	}
}

func (h *CollisionHandler) handleCollision(player *entity.Player, entity *entity.GameEntity) {
	if h.isHorizontalCollision(player, entity) {
		h.handleHorizontalCollision(player, entity)
		h.handleVerticalCollision(player, entity)
	} else {
		h.handleVerticalCollision(player, entity)
		h.handleHorizontalCollision(player, entity)
	}
}

func (h *CollisionHandler) isHorizontalCollision(player *entity.Player, entity *entity.GameEntity) bool {
	return player.Bounds().Intersect(entity.Bounds()).Dx() < player.Bounds().Intersect(entity.Bounds()).Dy()
}

func (h *CollisionHandler) handleVerticalCollision(player *entity.Player, entity *entity.GameEntity) {
	entityBounds := entity.Bounds()

	if h.playerCollidesOnTopOfEntity(player, entity) {
		player.SetPositionY(entityBounds.Min.Y - player.Height())
		player.Land()
		return
	}

	if h.playerCollidesOnBottomOfEntity(player, entity) {
		player.SetPositionY(entityBounds.Max.Y)
		player.Land()
	}
}

func (h *CollisionHandler) playerCollidesOnTopOfEntity(player *entity.Player, entity *entity.GameEntity) bool {
	if !player.Bounds().Overlaps(entity.Bounds()) {
		return false
	}

	playerBounds := player.Bounds()
	entityBounds := entity.Bounds()

	return player.VerticalVelocity() > 0 &&
		playerBounds.Min.Y < entityBounds.Min.Y &&
		playerBounds.Max.Y > entityBounds.Min.Y
}

func (h *CollisionHandler) playerCollidesOnBottomOfEntity(player *entity.Player, entity *entity.GameEntity) bool {
	if !player.Bounds().Overlaps(entity.Bounds()) {
		return false
	}

	playerBounds := player.Bounds()
	entityBounds := entity.Bounds()

	return player.VerticalVelocity() < 0 &&
		playerBounds.Min.Y > entityBounds.Min.Y &&
		playerBounds.Min.Y < entityBounds.Max.Y
}

func (h *CollisionHandler) handleHorizontalCollision(player *entity.Player, entity *entity.GameEntity) {
	entityBounds := entity.Bounds()

	if h.playerCollidesOnLeftOfEntity(player, entity) {
		player.SetPositionX(entityBounds.Min.X - player.Width())
		player.Stop()
		return
	}

	if h.playerCollidesOnRightOfEntity(player, entity) {
		player.SetPositionX(entityBounds.Max.X)
		player.Stop()
	}
}

func (h *CollisionHandler) playerCollidesOnLeftOfEntity(player *entity.Player, entity *entity.GameEntity) bool {
	if !player.Bounds().Overlaps(entity.Bounds()) {
		return false
	}

	playerBounds := player.Bounds()
	entityBounds := entity.Bounds()

	return player.HorizontalVelocity() > 0 &&
		playerBounds.Min.X < entityBounds.Min.X &&
		playerBounds.Max.X > entityBounds.Min.X
}

func (h *CollisionHandler) playerCollidesOnRightOfEntity(player *entity.Player, entity *entity.GameEntity) bool {
	if !player.Bounds().Overlaps(entity.Bounds()) {
		return false
	}

	playerBounds := player.Bounds()
	entityBounds := entity.Bounds()

	return player.HorizontalVelocity() < 0 &&
		playerBounds.Min.X < entityBounds.Max.X &&
		playerBounds.Max.X > entityBounds.Max.X
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
