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
	minX := e.radius
	maxX := screenWidth - e.radius
	e.x += e.vx
	if e.x <= minX || e.x >= maxX {
		if e.x <= minX {
			e.x = minX
		} else {
			e.x = maxX
		}
		e.vx *= -1
	}

	minY := e.radius
	maxY := screenHeight - e.radius
	e.y += e.vy
	if e.y <= minY || e.y >= maxY {
		if e.y <= minY {
			e.y = minY
		} else {
			e.y = maxY
		}

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
		if y1 < e.y {
			for y2 := y1; y2 <= e.y; y2++ {
				screen.Set(x1, y2, orange)
			}
		} else {
			for y2 := y1; y2 > e.y; y2-- {
				screen.Set(x1, y2, orange)
			}
		}
		screen.Set(x1, y1, orange)
	}
}
