package asteroids

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ship struct {
	Entity
}

func makeShip() (*Ship, error) {
	ship := &Ship{
		Entity: NewEntity(),
	}

	ship.Position.X = 100
	ship.Position.Y = 100

	img, _, err := ebitenutil.NewImageFromFile("src/assets/sprites/PLAYER.png")
	if err != nil {
		return nil, err
	}
	ship.Sprite = NewSprite(img)

	return ship, nil
}
