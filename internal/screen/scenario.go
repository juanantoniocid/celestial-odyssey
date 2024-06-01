package screen

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/world/entities"
)

const (
	sideMargin = 8
)

type InputHandler interface {
	UpdatePlayer()
}

type Renderer interface {
	DrawBackground(screen *ebiten.Image, background *ebiten.Image)
	DrawPlayer(screen *ebiten.Image, player *entities.Player)
	DrawCollidable(screen *ebiten.Image, collidable entities.Collidable)
}

type ScenarioImpl struct {
	player      *entities.Player
	collidables []entities.Collidable

	background   *ebiten.Image
	renderer     Renderer
	inputHandler InputHandler

	width  int
	height int
}

func NewScenario(player *entities.Player, background *ebiten.Image, renderer Renderer, inputHandler InputHandler, width int, height int) *ScenarioImpl {
	return &ScenarioImpl{
		player:       player,
		background:   background,
		renderer:     renderer,
		inputHandler: inputHandler,
		width:        width,
		height:       height,
		collidables:  make([]entities.Collidable, 0),
	}
}

func (s *ScenarioImpl) AddCollidable(c entities.Collidable) {
	s.collidables = append(s.collidables, c)
}

func (s *ScenarioImpl) Update() error {
	s.inputHandler.UpdatePlayer()
	s.player.Update()

	s.checkCollisions()
	s.checkIfPlayerIsOnPlatform()
	s.enforceBoundaries()

	return nil
}

func (s *ScenarioImpl) checkCollisions() {
	for _, c := range s.collidables {
		if s.player.Bounds().Overlaps(c.Bounds()) {
			s.handleCollision(c)
		}
	}
}

func (s *ScenarioImpl) handleCollision(c entities.Collidable) {
	playerBounds := s.player.Bounds()
	collidableBounds := c.Bounds()

	// Check for vertical collisions first (jumping and landing)
	if playerBounds.Min.Y < collidableBounds.Max.Y && playerBounds.Max.Y > collidableBounds.Min.Y {
		// Landing on top of the box
		if s.player.Position().Y < collidableBounds.Min.Y && playerBounds.Max.Y > collidableBounds.Min.Y {
			s.player.SetPositionY(collidableBounds.Min.Y - s.player.Height())
			s.player.Land()
			return
		}
		// Hitting the bottom of the box
		if s.player.Position().Y > collidableBounds.Max.Y && playerBounds.Min.Y < collidableBounds.Max.Y {
			s.player.SetPositionY(collidableBounds.Max.Y)
			s.player.Land()
			return
		}
	}

	// Horizontal collision
	if playerBounds.Min.X < collidableBounds.Max.X && playerBounds.Max.X > collidableBounds.Min.X {
		// Hitting the right side of the box
		if playerBounds.Min.X < collidableBounds.Min.X && playerBounds.Max.X > collidableBounds.Min.X {
			s.player.SetPositionX(collidableBounds.Min.X - s.player.Width())
			s.player.Stop()
			return
		}

		// Hitting the left side of the box
		if playerBounds.Min.X < collidableBounds.Max.X && playerBounds.Max.X > collidableBounds.Max.X {
			s.player.SetPositionX(collidableBounds.Max.X)
			s.player.Stop()
			return
		}
	}
}

func (s *ScenarioImpl) enforceBoundaries() {
	if s.player.Position().X < 0 {
		s.player.SetPositionX(0)
	} else if s.player.Position().X+s.player.Width() > s.width {
		s.player.SetPositionX(s.width - s.player.Width())
	}

	if s.player.Position().Y < 0 {
		s.player.SetPositionY(0)
	} else if s.player.Position().Y+s.player.Height() > s.height {
		s.player.SetPositionY(s.height - s.player.Height())
		s.player.Land()
	}
}

func (s *ScenarioImpl) checkIfPlayerIsOnPlatform() {
	playerBounds := s.player.Bounds()
	isOnPlatform := false

	// Check if the player is on any platform
	for _, c := range s.collidables {
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
	if playerBounds.Max.Y >= s.height {
		isOnPlatform = true
	}

	// If the player is not on any platform and not jumping, apply gravity
	if !isOnPlatform && !s.player.IsJumping() {
		s.player.Fall()
	}
}

func (s *ScenarioImpl) Draw(screen *ebiten.Image) {
	s.renderer.DrawBackground(screen, s.background)
	s.renderer.DrawPlayer(screen, s.player)
	for _, c := range s.collidables {
		s.renderer.DrawCollidable(screen, c)
	}
}

func (s *ScenarioImpl) ShouldTransitionRight() bool {
	return s.player.Position().X+s.player.Width() >= s.width
}

func (s *ScenarioImpl) ShouldTransitionLeft() bool {
	return s.player.Position().X <= 0
}

func (s *ScenarioImpl) SetPlayerPositionAtLeft() {
	s.player.SetPositionX(sideMargin)
}

func (s *ScenarioImpl) SetPlayerPositionAtRight() {
	s.player.SetPositionX(s.width - sideMargin - s.player.Width())
}
