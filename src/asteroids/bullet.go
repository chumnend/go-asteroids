package asteroids

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	BULLET_SPEED = 5
)

type Bullet struct {
	Entity
}

func makeBullet() (*Bullet, error) {
	bullet := &Bullet{
		Entity: NewEntity(),
	}

	// make sure bullets are hidden initially
	bullet.IsHidden = true

	// load the player sprite
	img, _, err := ebitenutil.NewImageFromFile("src/assets/sprites/BULLET.png")
	if err != nil {
		return nil, err
	}
	bullet.Sprite = NewSprite(img)

	return bullet, nil
}

func (b *Bullet) show() {
	b.IsHidden = false
}

func (b *Bullet) setPosition(x, y int) {
	b.X = x
	b.Y = y
}

func (b *Bullet) setDirection(dir int) {
	b.Vx = int(float64(BULLET_SPEED) * math.Sin(float64(dir)))
	b.Vy = int(float64(BULLET_SPEED) * math.Cos(float64(dir)))
	fmt.Printf("Vx: %v, Vy: %v\n", b.Vx, b.Vy)
}

func (b *Bullet) updatePosition() {
	b.X += b.Vx
	b.Y -= b.Vy
}
