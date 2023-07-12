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

	// get random position
	randX := rand.Intn(gameWidth - img.Bounds().Dy())
	randY := rand.Intn(gameHeight - img.Bounds().Dy())
	asteroid.Position.X = randX
	asteroid.Position.Y = randY

	// set velocity
	asteroid.Velocity.Vx = 1
	asteroid.Velocity.Vy = -1

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

	a.Position.X += a.Velocity.Vx
	if float64(a.Position.X)+rect.W >= gameWidth || float64(a.Position.X) <= 0 {
		a.Velocity.Vx *= -1
	}

	a.Position.Y += a.Velocity.Vy
	if float64(a.Position.Y)+rect.H >= gameHeight || float64(a.Position.Y) <= 0 {
		a.Velocity.Vy *= -1
	}
}
