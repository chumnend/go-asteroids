package asteroids

import (
	"fmt"
	"image"
	"math"
)

type Vec2f struct {
	X, Y float64
}

func (v Vec2f) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", v.X, v.Y)
}

type FloatRect struct {
	X, Y, W, H float64
}

func (r FloatRect) ToImageRect() image.Rectangle {
	return image.Rect(int(r.X), int(r.Y), int(r.X+r.W), int(r.Y+r.H))
}

func (r FloatRect) Corners() (float64, float64, float64, float64) {
	return r.X, r.Y, r.X + r.W, r.Y + r.H
}

func degreeToRad(deg float64) float64 {
	return deg * math.Pi / 180
}

func radToDegree(rad float64) float64 {
	return rad * 180 / math.Pi
}
