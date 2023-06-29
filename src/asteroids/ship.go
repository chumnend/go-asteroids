package asteroids

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ship struct {
	Position
}

func makeShip() (*Ship, error) {
	return &Ship{}, nil
}

// Draw takes an ebiten screen object and draws the sprite on it (PLACEHOLDER)
func (s *Ship) Draw(screen *ebiten.Image) {
	purple := color.RGBA{255, 0, 255, 255}
	for x := s.X; x < s.X+50; x++ {
		for y := s.Y; y < s.Y+50; y++ {
			screen.Set(x, y, purple)
		}
	}
}
