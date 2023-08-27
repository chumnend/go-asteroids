package asteroids

import (
	"math/rand"
)

const (
	INITIAL_NUMBER_OF_ASTEROIDS = 3
	INITIAL_ASTEROID_VX         = 1
	INITIAL_ASTEROID_VY         = -1
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
	asteroid.sprite, err = NewSpriteFromImagePath("src/assets/sprites/rock.png")
	if err != nil {
		return nil, err
	}

	// set initial asteroid parameters
	asteroid.Initialize()

	return asteroid, nil
}

func (asteroid *Asteroid) GetRandomPosition() {
	randX := rand.Intn(GAME_WIDTH - int(asteroid.sprite.GetSize().X))
	randY := rand.Intn(GAME_HEIGHT - int(asteroid.sprite.GetSize().Y))
	asteroid.position.X = float64(randX)
	asteroid.position.Y = float64(randY)
}

func (asteroid *Asteroid) Initialize() {
	asteroid.GetRandomPosition()
	asteroid.velocity.X = INITIAL_ASTEROID_VX
	asteroid.velocity.Y = INITIAL_ASTEROID_VY

}

func (asteroid *Asteroid) Update() {
	rect := asteroid.GetBoundingRect()

	asteroid.position.X += asteroid.velocity.X
	if asteroid.position.X+rect.W >= GAME_WIDTH || asteroid.position.X <= 0 {
		asteroid.velocity.X *= -1
	}

	asteroid.position.Y += asteroid.velocity.Y
	if asteroid.position.Y+rect.H >= GAME_HEIGHT || asteroid.position.Y <= 0 {
		asteroid.velocity.Y *= -1
	}
}
