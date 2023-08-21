package asteroids

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	X         int
	Y         int
	Direction int // in degrees
	Vx        int // pixel per draw
	Vy        int // pixel per draw
	Sprite    *Sprite
	IsHidden  bool
}

func NewEntity() Entity {
	return Entity{
		Sprite: &Sprite{},
	}
}

func (e *Entity) getAABB() FloatRect {
	w := e.Sprite.GetSize().X
	h := e.Sprite.GetSize().Y
	return FloatRect{float64(e.X), float64(e.Y), w, h}
}

func (e *Entity) Draw(screen *ebiten.Image) {
	// if object is hidden do not draw
	if e.IsHidden {
		return
	}

	var m ebiten.GeoM

	// rotate the image into correct direction
	m.Rotate(degreeToRad(float64(e.Direction)))

	// move image to correct location
	m.Translate(
		float64(e.X),
		float64(e.Y),
	)
	e.Sprite.Draw(screen, m)
}
