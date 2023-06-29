package asteroids

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Asteroid struct {
	Position
}

func makeAsteroid() (*Asteroid, error) {
	return &Asteroid{}, nil
}

func (o *Asteroid) Draw(screen *ebiten.Image) {
	orange := color.RGBA{R: 255, G: 127, B: 80}
	minAngle := math.Acos(1 - 1/float64(10))
	for angle := float64(0); angle <= 360; angle += minAngle {
		xDelta := float64(10) * math.Cos(angle)
		yDelta := float64(10) * math.Sin(angle)
		x1 := int(math.Round(float64(o.X) + xDelta))
		y1 := int(math.Round(float64(o.Y) + yDelta))
		if y1 < o.Y {
			for y2 := y1; y2 <= o.Y; y2++ {
				screen.Set(x1, y2, orange)
			}
		} else {
			for y2 := y1; y2 > o.Y; y2-- {
				screen.Set(x1, y2, orange)
			}
		}
		screen.Set(x1, y1, orange)
	}
}
