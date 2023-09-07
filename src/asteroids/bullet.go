package asteroids

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	NUMBER_OF_BULLETS = 5
	BULLET_SPEED      = 4
)

type Bullet struct {
	Entity
}

func NewBullet() (*Bullet, error) {
	bullet := &Bullet{
		Entity: NewEntity(),
	}

	// load the bullet sprite
	var err error
	bullet.Sprite, err = NewSpriteFromImagePath("src/assets/sprites/bullet.png")
	if err != nil {
		return nil, err
	}

	// set initial ship parameters
	bullet.Initialize()

	return bullet, nil
}

func (bullet *Bullet) Shoot(shipX, shipY, shipDir float64) {
	bullet.Position.X = shipX
	bullet.Position.Y = shipY
	bullet.Direction = shipDir
	bullet.IsHidden = false

	dirRad := degreeToRad(shipDir - 90) // assuming direction 0 deg is UP, need to rotate -ve
	bullet.Velocity.X = BULLET_SPEED * math.Cos(dirRad)
	bullet.Velocity.Y = BULLET_SPEED * math.Sin(dirRad)
}

func (bullet *Bullet) Initialize() {
	bullet.IsHidden = true
	bullet.Position.X = GAME_HEIGHT * 2
	bullet.Position.Y = GAME_WIDTH * 2
	bullet.Velocity.X = 0
	bullet.Velocity.Y = 0
}

func (bullet *Bullet) Update() {
	bullet.Position.X += bullet.Velocity.X
	bullet.Position.Y += bullet.Velocity.Y
}

type Bullets []*Bullet

var currentBulletIdx int

func NewBullets() (Bullets, error) {
	bullets := make(Bullets, 0)
	for i := 0; i < NUMBER_OF_BULLETS; i++ {
		bullet, err := NewBullet()
		if err != nil {
			return nil, err
		}
		bullets = append(bullets, bullet)
	}

	return bullets, nil
}

func (bullets Bullets) Shoot(ship *Ship) {
	// if out of bullets, do not shoot
	if currentBulletIdx >= NUMBER_OF_BULLETS {
		return
	}
	bullets[currentBulletIdx].Shoot(ship.Position.X, ship.Position.Y, ship.Direction)
	currentBulletIdx += 1
}

func (bullets Bullets) Initialize() {
	currentBulletIdx = 0
	for _, bullet := range bullets {
		bullet.Initialize()
	}
}

func (bullets Bullets) Update() {
	for _, bullet := range bullets {
		bullet.Update()
	}
}

func (bullets Bullets) Draw(screen *ebiten.Image) {
	for _, bullet := range bullets {
		bullet.Draw(screen)
	}
}
