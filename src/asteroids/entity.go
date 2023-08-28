package asteroids

import "github.com/hajimehoshi/ebiten/v2"

type Entity struct {
	position    Vector2
	velocity    Vector2
	direction   float64 // in degrees
	sprite      *Sprite
	isHidden    bool
	isRotatable bool
}

func NewEntity() Entity {
	return Entity{}
}

func (e *Entity) GetBoundingRect() Rectangle {
	w := e.sprite.GetSize().X
	h := e.sprite.GetSize().Y
	return Rectangle{float64(e.position.X), float64(e.position.Y), w, h}
}

func (e *Entity) CollidesWith(obj *Entity) bool {
	rect1 := e.GetBoundingRect().ToImageRect()
	rect2 := obj.GetBoundingRect().ToImageRect()
	return rect1.Overlaps(rect2)
}

func (e *Entity) Draw(screen *ebiten.Image) {
	// if object is hidden do not draw
	if e.isHidden {
		return
	}

	var m ebiten.GeoM

	// rotate the image into correct direction
	if e.isRotatable {
		m.Rotate(degreeToRad(float64(e.direction)))
	}

	// move image to correct location
	m.Translate(
		float64(e.position.X),
		float64(e.position.Y),
	)
	e.sprite.Draw(screen, m)
}
