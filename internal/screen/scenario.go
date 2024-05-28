package screen

import (
	entities2 "celestial-odyssey/internal/world/entities"
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/graphics"
)

const (
	sideMargin = 8
)

type ScenarioImpl struct {
	player     *entities2.Player
	background *ebiten.Image
	renderer   *graphics.Renderer

	width  int
	height int

	collidables []entities2.Collidable
}

func NewScenario(player *entities2.Player, background *ebiten.Image, renderer *graphics.Renderer, width, height int) *ScenarioImpl {
	return &ScenarioImpl{
		player:      player,
		background:  background,
		renderer:    renderer,
		width:       width,
		height:      height,
		collidables: make([]entities2.Collidable, 0),
	}
}

func (s *ScenarioImpl) AddCollidable(c entities2.Collidable) {
	s.collidables = append(s.collidables, c)
}

func (s *ScenarioImpl) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		s.player.MoveLeft()
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		s.player.MoveRight()
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		s.player.Jump()
	}

	s.player.Update()
	s.checkCollisions()

	return nil
}

func (s *ScenarioImpl) checkCollisions() {
	for _, c := range s.collidables {
		if s.player.Bounds().Overlaps(c.Bounds()) {
			s.handleCollision(c)
		}
	}
}

func (s *ScenarioImpl) handleCollision(c entities2.Collidable) {
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
			s.player.VelocityY = 0
			return
		}
	}

	// Check for horizontal collisions (moving left and right)
	if playerBounds.Min.X < collidableBounds.Max.X && playerBounds.Max.X > collidableBounds.Min.X {
		// Hitting the right side of the box
		if s.player.Position().X < collidableBounds.Min.X && playerBounds.Max.X > collidableBounds.Min.X {
			s.player.SetPositionX(collidableBounds.Min.X - s.player.Width())
			s.player.Stop()
			return
		}
		// Hitting the left side of the box
		if s.player.Position().X > collidableBounds.Max.X && playerBounds.Min.X < collidableBounds.Max.X {
			s.player.SetPositionX(collidableBounds.Max.X)
			s.player.Stop()
			return
		}
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
	s.player.SetPositionAtBottomLeft(sideMargin)
}

func (s *ScenarioImpl) SetPlayerPositionAtRight() {
	s.player.SetPositionAtBottomRight(s.width - sideMargin)
}
