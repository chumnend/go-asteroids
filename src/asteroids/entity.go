package asteroids

import "github.com/hajimehoshi/ebiten/v2"

type Entity struct {
	Position    Vector2
	Velocity    Vector2
	Direction   float64 // in degrees, we assume 0 is up and entity is placed facing up
	Sprite      *Sprite
	IsHidden    bool
	IsRotatable bool
}

func NewEntity() Entity {
	return Entity{
		Position:    Vector2{X: 0, Y: 0},
		Velocity:    Vector2{X: 0, Y: 0},
		Direction:   0,
		Sprite:      nil,
		IsHidden:    false,
		IsRotatable: false,
	}
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

func (e *Entity) CollidesWithOffset(obj *Entity, offset float64) bool {
	w := e.Sprite.GetSize().X
	h := e.Sprite.GetSize().Y
	rect1 := Rectangle{float64(e.Position.X - offset), float64(e.Position.Y - offset), w + offset, h + offset}.ToImageRect()
	rect2 := obj.GetBoundingRect().ToImageRect()
	return rect1.Overlaps(rect2)
}

func (e *Entity) Draw(screen *ebiten.Image) {
	// if object is hidden do not draw
	if e.IsHidden {
		return
	}

	var m ebiten.GeoM

	// rotate the image if image is rotable to match Direction
	// the origin is normally in the upper left of the sprite, to rotate properly
	// we much first update the origin using the sprites dimensions
	if e.IsRotatable {
		originX := -float64(e.Sprite.GetSize().X) / 2
		originY := -float64(e.Sprite.GetSize().Y) / 2
		m.Translate(originX, originY)
		m.Rotate(degreeToRad(float64(e.Direction)))
	}

	// move image to correct location
	m.Translate(
		float64(e.Position.X),
		float64(e.Position.Y),
	)
	e.Sprite.Draw(screen, m)
}
