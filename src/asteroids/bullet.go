package asteroids

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	NUMBER_OF_BULLETS = 3
	BULLET_SPEED      = 4
)

type Bullet struct {
	Entity
	IsFree bool
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

func (bullet *Bullet) Shoot(ship *Ship) {
	bullet.IsHidden = false
	bullet.IsFree = false
	bullet.Position.X = ship.Position.X
	bullet.Position.Y = ship.Position.Y
	bullet.Direction = ship.Direction

	dirRad := degreeToRad(ship.Direction - 90) // assuming direction 0 deg is UP, need to rotate -ve
	bullet.Velocity.X = BULLET_SPEED * math.Cos(dirRad)
	bullet.Velocity.Y = BULLET_SPEED * math.Sin(dirRad)
}

func (bullet *Bullet) Initialize() {
	bullet.IsHidden = true
	bullet.IsFree = true
	bullet.Position.X = GAME_HEIGHT * 2
	bullet.Position.Y = GAME_WIDTH * 2
	bullet.Velocity.X = 0
	bullet.Velocity.Y = 0
}

func (bullet *Bullet) Update() {
	bullet.Position.X += bullet.Velocity.X
	bullet.Position.Y += bullet.Velocity.Y

	// if bullet off screen, reset the bullet
	if bullet.Position.X < 0 || bullet.Position.X+bullet.Sprite.GetSize().X > GAME_WIDTH || bullet.Position.Y < 0 || bullet.Position.Y+bullet.Sprite.GetSize().Y > GAME_HEIGHT {
		bullet.Initialize()
	}
}

type Bullets []*Bullet

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
	// look for free bullet to fire
	currentIdx := 0
	for currentIdx < NUMBER_OF_BULLETS {
		if bullets[currentIdx].IsFree {
			bullets[currentIdx].Shoot(ship)
			return
		}
		currentIdx++
	}
}

func (bullets Bullets) Initialize() {
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
