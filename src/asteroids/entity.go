package asteroids

import "github.com/hajimehoshi/ebiten/v2"

type Entity struct {
	Position    Vector2
	Velocity    Vector2
	Direction   float64 // in degrees
	Sprite      *Sprite
	IsHidden    bool
	IsRotatable bool
}

func NewEntity() Entity {
	return Entity{}
}

func (e *Entity) GetBoundingRect() Rectangle {
	w := e.Sprite.GetSize().X
	h := e.Sprite.GetSize().Y
	return Rectangle{float64(e.Position.X), float64(e.Position.Y), w, h}
}

func (e *Entity) CollidesWith(obj *Entity) bool {
	rect1 := e.GetBoundingRect().ToImageRect()
	rect2 := obj.GetBoundingRect().ToImageRect()
	return rect1.Overlaps(rect2)
}

func (e *Entity) Draw(screen *ebiten.Image) {
	// if object is hidden do not draw
	if e.IsHidden {
		return
	}

	var m ebiten.GeoM

	// rotate the image into correct Direction
	if e.IsRotatable {
		m.Rotate(degreeToRad(float64(e.Direction)))
	}

	// move image to correct location
	m.Translate(
		float64(e.Position.X),
		float64(e.Position.Y),
	)
	e.Sprite.Draw(screen, m)
}
