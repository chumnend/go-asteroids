package asteroids

import "math"

const (
	BULLET_SPEED = 4
)

type Bullet struct {
	Entity
}

func NewBullet() (*Bullet, error) {
	bullet := &Bullet{
		Entity: NewEntity(),
	}

	// load the asteroid sprite
	var err error
	bullet.Sprite, err = NewSpriteFromImagePath("src/assets/sprites/bullet.png")
	if err != nil {
		return nil, err
	}

	// set initial ship parameters
	bullet.Initialize()

	return bullet, nil
}

func (bullet *Bullet) Fire(shipX, shipY, shipDir float64) {
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
}

func (bullet *Bullet) Update() {
	bullet.Position.X += bullet.Velocity.X
	bullet.Position.Y += bullet.Velocity.Y
}
