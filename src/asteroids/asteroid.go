package asteroids

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	INITIAL_NUMBER_OF_ASTEROIDS = 3
	INITIAL_ASTEROID_VX         = 1
	INITIAL_ASTEROID_VY         = -1
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
	// TODO: dont spawn asteroid on ship
	randX := rand.Intn(gameWidth - img.Bounds().Dy())
	randY := rand.Intn(gameHeight - img.Bounds().Dy())
	asteroid.X = randX
	asteroid.Y = randY

	// set velocity
	asteroid.Vx = INITIAL_ASTEROID_VX
	asteroid.Vy = INITIAL_ASTEROID_VY

	return asteroid, nil
}

func makeAsteroids() ([]*Asteroid, error) {
	numberOfAsteroids := INITIAL_NUMBER_OF_ASTEROIDS
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

func (a *Asteroid) resetPosition() {
	// TODO: dont spawn asteroid on ship
	randX := rand.Intn(gameWidth - a.Sprite.Image.Bounds().Dy())
	randY := rand.Intn(gameHeight - a.Sprite.Image.Bounds().Dy())
	a.X = randX
	a.Y = randY
}

func (a *Asteroid) bounce() {
	a.Vx *= -1
	a.Vy *= -1
}
