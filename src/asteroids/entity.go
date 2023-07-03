package asteroids

import "github.com/hajimehoshi/ebiten/v2"

type Entity struct {
	Position *Position
	Velocity *Velocity
	Sprite   *Sprite
}

func NewEntity() Entity {
	return Entity{
		Position: &Position{},
		Velocity: &Velocity{},
		Sprite:   &Sprite{},
	}
}

func (e *Entity) getAABB() FloatRect {
	w := e.Sprite.GetSize().X
	h := e.Sprite.GetSize().Y
	return FloatRect{float64(e.Position.X), float64(e.Position.Y), w, h}
}

func (e *Entity) Draw(screen *ebiten.Image) {
	var m ebiten.GeoM
	m.Translate(
		float64(e.Position.X),
		float64(e.Position.Y),
	)
	e.Sprite.Draw(screen, m)
}
