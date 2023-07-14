package asteroids

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Asteroid struct {
	Entity
}

func makeAsteroid() (*Asteroid, error) {
	asteroid := &Asteroid{
		Entity: NewEntity(),
	}

	// load asteroid sprite
	img, _, err := ebitenutil.NewImageFromFile("src/assets/sprites/ROCK.png")
	if err != nil {
		return nil, err
	}
	asteroid.Sprite = NewSprite(img)

	// get random initial position
	randX := rand.Intn(gameWidth - img.Bounds().Dy())
	randY := rand.Intn(gameHeight - img.Bounds().Dy())
	asteroid.X = randX
	asteroid.Y = randY

	// set velocity
	asteroid.Vx = 1
	asteroid.Vy = -1

	return asteroid, nil
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

func (a *Asteroid) updatePosition() {
	rect := a.getAABB()

	a.X += a.Vx
	if float64(a.X)+rect.W >= gameWidth || float64(a.X) <= 0 {
		a.Vx *= -1
	}

	a.Y += a.Vy
	if float64(a.Y)+rect.H >= gameHeight || float64(a.Y) <= 0 {
		a.Vy *= -1
	}
}
