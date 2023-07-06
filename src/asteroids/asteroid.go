package asteroids

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Asteroid struct {
	Entity
}

func makeAsteroids() ([]*Asteroid, error) {
	numberOfAsteroids := 3
	asteroids := make([]*Asteroid, 0)

	for i := 0; i < numberOfAsteroids; i++ {
		asteroid, err := makeAsteroid()
		if err != nil {
			return nil, err
		}
		asteroids = append(asteroids, asteroid)
	}

	return asteroids, nil
}

func makeAsteroid() (*Asteroid, error) {
	asteroid := &Asteroid{
		Entity: NewEntity(),
	}

	// get random position
	randX := rand.Intn(gameWidth)
	randY := rand.Intn(gameHeight)
	asteroid.Position.X = randX
	asteroid.Position.Y = randY

	// load asteroid sprite
	img, _, err := ebitenutil.NewImageFromFile("src/assets/sprites/ROCK.png")
	if err != nil {
		return nil, err
	}
	asteroid.Sprite = NewSprite(img)

	return asteroid, nil
}
