package asteroids

import "github.com/hajimehoshi/ebiten/v2"

type Entity struct {
	X         int
	Y         int
	Direction int

	Vx int
	Vy int

	Sprite *Sprite
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
	var m ebiten.GeoM
	m.Translate(
		float64(e.X),
		float64(e.Y),
	)
	e.Sprite.Draw(screen, m)
}
