package gunthur

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	x      int
	y      int
	vx     int
	vy     int
	radius int
}

// NewEnemy returns an Enemy struct
func NewEnemy() *Enemy {
	return &Enemy{
		x:      25,
		y:      50,
		vx:     1,
		vy:     1,
		radius: 15,
	}
}

func (e *Enemy) Update(keys []ebiten.Key) error {
	e.x += e.vx
	e.y += e.vy
	if e.x-e.radius <= 0 || e.x+e.radius >= screenWidth {
		e.vx *= -1
	}
	if e.y-e.radius <= 0 || e.y+e.radius >= screenHeight {
		e.vy *= -1
	}

	return nil
}

func (e *Enemy) Draw(screen *ebiten.Image, opts ebiten.DrawImageOptions) {
	orange := color.RGBA{R: 255, G: 127, B: 80}
	minAngle := math.Acos(1 - 1/float64(e.radius))
	for angle := float64(0); angle <= 360; angle += minAngle {
		xDelta := float64(e.radius) * math.Cos(angle)
		yDelta := float64(e.radius) * math.Sin(angle)
		x1 := int(math.Round(float64(e.x) + xDelta))
		y1 := int(math.Round(float64(e.y) + yDelta))
		screen.Set(x1, y1, orange)
	}
}
