package asteroids

import "github.com/hajimehoshi/ebiten/v2/ebitenutil"

type Bullet struct {
	Entity
}

func makeBullet() (*Bullet, error) {
	bullet := &Bullet{
		Entity: NewEntity(),
	}

	// set initial bullet position
	bullet.X = 0
	bullet.Y = 0
	bullet.Vx = 0
	bullet.Vy = 1

	// load the player sprite
	img, _, err := ebitenutil.NewImageFromFile("src/assets/sprites/BULLET.png")
	if err != nil {
		return nil, err
	}
	bullet.Sprite = NewSprite(img)

	return bullet, nil
}
