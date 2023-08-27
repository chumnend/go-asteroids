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
	asteroid.GetRandomPosition()
	asteroid.velocity.X = INITIAL_ASTEROID_VX
	asteroid.velocity.Y = INITIAL_ASTEROID_VY

	return asteroid, nil
}

func (asteroid *Asteroid) GetRandomPosition() {
	randX := rand.Intn(GAME_WIDTH - int(asteroid.sprite.GetSize().X))
	randY := rand.Intn(GAME_HEIGHT - int(asteroid.sprite.GetSize().Y))
	asteroid.position.X = float64(randX)
	asteroid.position.Y = float64(randY)
}
