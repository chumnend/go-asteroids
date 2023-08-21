package asteroids

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	SHIP_START_X      = gameWidth / 2
	SHIP_START_Y      = gameHeight / 2
	INITIAL_DIRECTION = 0 // degrees
	INITIAL_SHIP_VX   = 2
	INITIAL_SHIP_VY   = 2
	TURN_RATE         = 5
)

type Ship struct {
	Entity
}

func makeShip() (*Ship, error) {
	ship := &Ship{
		Entity: NewEntity(),
	}

	// set initial player position
	ship.X = SHIP_START_X
	ship.Y = SHIP_START_Y
	ship.Direction = INITIAL_DIRECTION
	ship.Vx = INITIAL_SHIP_VX
	ship.Vy = INITIAL_SHIP_VY

	// load the player sprite
	img, _, err := ebitenutil.NewImageFromFile("src/assets/sprites/PLAYER.png")
	if err != nil {
		return nil, err
	}
	ship.Sprite = NewSprite(img)

	return ship, nil
}

func (s *Ship) resetPosition() {
	s.X = SHIP_START_X
	s.Y = SHIP_START_Y
}

func (s *Ship) moveUp() {
	newPos := s.Y - s.Vy
	if newPos > 0 {
		s.Y = newPos
	}
}

func (s *Ship) moveDown() {
	rect := s.getAABB()
	newPos := s.Y + s.Vy
	if newPos+int(rect.H) < gameHeight {
		s.Y = newPos
	}
}

func (s *Ship) moveLeft() {
	newPos := s.X - s.Vx
	if newPos > 0 {
		s.X = newPos
	}
}

func (s *Ship) moveRight() {
	rect := s.getAABB()
	newPos := s.X + s.Vx
	if newPos+int(rect.W) < gameWidth {
		s.X = newPos
	}
}

func (s *Ship) rotateLeft() {
	s.Direction -= TURN_RATE
}

func (s *Ship) rotateRight() {
	s.Direction += TURN_RATE
}
