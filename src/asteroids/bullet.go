package asteroids

import "github.com/hajimehoshi/ebiten/v2/ebitenutil"

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

func (b *Bullet) updatePosition() {
	b.X += b.Vx
	b.Y -= b.Vy
}
