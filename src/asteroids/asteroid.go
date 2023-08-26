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

	// get random initial position
	// TODO: dont spawn asteroid on ship
	randX := rand.Intn(GAME_WIDTH - int(asteroid.sprite.GetSize().X))
	randY := rand.Intn(GAME_HEIGHT - int(asteroid.sprite.GetSize().Y))
	asteroid.position.X = float64(randX)
	asteroid.position.Y = float64(randY)

	// set velocity
	asteroid.velocity.X = INITIAL_ASTEROID_VX
	asteroid.velocity.Y = INITIAL_ASTEROID_VY

	return asteroid, nil
}
