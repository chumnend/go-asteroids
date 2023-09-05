package asteroids

import (
	"math/rand"
)

const (
	NUMBER_OF_ASTEROIDS = 5
	INITIAL_ASTEROID_VX = 1
	INITIAL_ASTEROID_VY = -1
)

type Asteroid struct {
	Entity
}

func NewAsteroid() (*Asteroid, error) {
	asteroid := &Asteroid{
		Entity: NewEntity(),
	}

	// load the asteroid sprite
	var err error
	asteroid.Sprite, err = NewSpriteFromImagePath("src/assets/sprites/rock.png")
	if err != nil {
		return nil, err
	}

	// set initial asteroid parameters
	asteroid.Initialize()

	return asteroid, nil
}

func NewAsteroids() ([]*Asteroid, error) {
	asteroids := make([]*Asteroid, 0)
	for i := 0; i < NUMBER_OF_ASTEROIDS; i++ {
		asteroid, err := NewAsteroid()
		if err != nil {
			return nil, err
		}
		asteroids = append(asteroids, asteroid)
	}

	return asteroids, nil
}

func (asteroid *Asteroid) GetRandomPosition() {
	randX := rand.Intn(GAME_WIDTH - int(asteroid.Sprite.GetSize().X))
	randY := rand.Intn(GAME_HEIGHT - int(asteroid.Sprite.GetSize().Y))
	asteroid.Position.X = float64(randX)
	asteroid.Position.Y = float64(randY)
}

func (asteroid *Asteroid) Bounce() {
	asteroid.Velocity.X *= -1
	asteroid.Velocity.Y *= -1
}

func (asteroid *Asteroid) Destroy() {
	asteroid.IsHidden = true
	asteroid.Position.X = GAME_HEIGHT * 2
	asteroid.Position.Y = GAME_WIDTH * 2
	asteroid.Velocity.X = 0
	asteroid.Velocity.Y = 0
}

func (asteroid *Asteroid) Initialize() {
	asteroid.IsHidden = false
	asteroid.GetRandomPosition()
	asteroid.Velocity.X = INITIAL_ASTEROID_VX
	asteroid.Velocity.Y = INITIAL_ASTEROID_VY
}

func (asteroid *Asteroid) Update() {
	rect := asteroid.GetBoundingRect()

	asteroid.Position.X += asteroid.Velocity.X
	if asteroid.Position.X+rect.W >= GAME_WIDTH || asteroid.Position.X <= 0 {
		asteroid.Velocity.X *= -1
	}

	asteroid.Position.Y += asteroid.Velocity.Y
	if asteroid.Position.Y+rect.H >= GAME_HEIGHT || asteroid.Position.Y <= 0 {
		asteroid.Velocity.Y *= -1
	}
}
