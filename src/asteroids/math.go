package asteroids

import (
	"image"
	"math"
)

type Vector2 struct {
	X, Y float64
}

func (v1 Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

type Rectangle struct {
	X, Y, W, H float64
}

func (rect Rectangle) ToImageRect() image.Rectangle {
	return image.Rect(int(rect.X), int(rect.Y), int(rect.X+rect.W), int(rect.Y+rect.H))
}

func (rect Rectangle) Corners() (float64, float64, float64, float64) {
	return rect.X, rect.Y, rect.X + rect.W, rect.Y + rect.H
}

func degreeToRad(deg float64) float64 {
	return deg * math.Pi / 180
}

func radToDegree(rad float64) float64 {
	return rad * 180 / math.Pi
}
