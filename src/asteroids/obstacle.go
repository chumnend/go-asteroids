package asteroids

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Obstacle struct {
	x      int
	y      int
	vx     int
	vy     int
	radius int
}

// NewObstacle returns an Obstacle struct
func NewObstacle() *Obstacle {
	return &Obstacle{
		x:      25,
		y:      50,
		vx:     1,
		vy:     1,
		radius: 15,
	}
}

func (o *Obstacle) Update(keys []ebiten.Key) error {
	minX := o.radius
	maxX := screenWidth - o.radius
	o.x += o.vx
	if o.x <= minX || o.x >= maxX {
		if o.x <= minX {
			o.x = minX
		} else {
			o.x = maxX
		}
		o.vx *= -1
	}

	minY := o.radius
	maxY := screenHeight - o.radius
	o.y += o.vy
	if o.y <= minY || o.y >= maxY {
		if o.y <= minY {
			o.y = minY
		} else {
			o.y = maxY
		}

		o.vy *= -1
	}

	return nil
}

func (o *Obstacle) Draw(screen *ebiten.Image, opts ebiten.DrawImageOptions) {
	orange := color.RGBA{R: 255, G: 127, B: 80}
	minAngle := math.Acos(1 - 1/float64(o.radius))
	for angle := float64(0); angle <= 360; angle += minAngle {
		xDelta := float64(o.radius) * math.Cos(angle)
		yDelta := float64(o.radius) * math.Sin(angle)
		x1 := int(math.Round(float64(o.x) + xDelta))
		y1 := int(math.Round(float64(o.y) + yDelta))
		if y1 < o.y {
			for y2 := y1; y2 <= o.y; y2++ {
				screen.Set(x1, y2, orange)
			}
		} else {
			for y2 := y1; y2 > o.y; y2-- {
				screen.Set(x1, y2, orange)
			}
		}
		screen.Set(x1, y1, orange)
	}
}
