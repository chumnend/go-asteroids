package asteroids

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	NUMBER_OF_ASTEROIDS = 5
	ASTEROID_INITIAL_VX = 1
	ASTEROID_INITIAL_VY = -1
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
	asteroid.Velocity.X = ASTEROID_INITIAL_VX
	asteroid.Velocity.Y = ASTEROID_INITIAL_VY
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

type Asteroids []*Asteroid

func NewAsteroids(ship *Ship) (Asteroids, error) {
	asteroids := make(Asteroids, 0)
	for i := 0; i < NUMBER_OF_ASTEROIDS; i++ {
		asteroid, err := NewAsteroid()
		if err != nil {
			return nil, err
		}
		asteroids = append(asteroids, asteroid)
	}

	// set initial parameters
	asteroids.Initialize(ship)

	return asteroids, nil
}

func (asteroids Asteroids) Initialize(ship *Ship) {
	// check to make sure asteroids do not spawn on ship, other asteroids
	for idx, asteroid := range asteroids {
		for asteroid.CollidesWith(&ship.Entity) {
			asteroid.GetRandomPosition()
		}

		otherAsteroids := make([]*Asteroid, len(asteroids))
		copy(otherAsteroids, asteroids)
		otherAsteroids = append(otherAsteroids[:idx], otherAsteroids[idx+1:]...)
		for _, oA := range otherAsteroids {
			for asteroid.CollidesWith(&oA.Entity) {
				asteroid.GetRandomPosition()
			}
		}
	}

	for _, asteroid := range asteroids {
		asteroid.Initialize()
	}
}

func (asteroids Asteroids) Update() {
	for _, asteroid := range asteroids {
		asteroid.Update()
	}
}

func (asteroids Asteroids) Draw(screen *ebiten.Image) {
	for _, asteroid := range asteroids {
		asteroid.Draw(screen)
	}
}
