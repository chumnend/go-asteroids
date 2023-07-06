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

	// set initail player position
	ship.Position.X = gameWidth / 2
	ship.Position.Y = gameHeight / 2

	// load the player sprite
	img, _, err := ebitenutil.NewImageFromFile("src/assets/sprites/PLAYER.png")
	if err != nil {
		return nil, err
	}
	ship.Sprite = NewSprite(img)

	return ship, nil
}
