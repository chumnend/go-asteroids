package gunthur

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Sprite represents a game sprite
type Sprite struct {
	x      int
	y      int
	width  int
	height int
}

// NewSprite returns a Sprite struct
func NewSprite() *Sprite {
	return &Sprite{
		x:      200,
		y:      200,
		width:  50,
		height: 50,
	}
}

// MoveTo changes the position of the sprite
func (s *Sprite) MoveTo(x int, y int) {
	s.x = x
	s.y = y
}

// Draw takes an ebiten screen object and draws the sprite on it
func (s *Sprite) Draw(screen *ebiten.Image) {
	purple := color.RGBA{255, 0, 255, 255}
	for x := s.x; x < s.x+s.width; x++ {
		for y := s.y; y < s.y+s.height; y++ {
			screen.Set(x, y, purple)
		}
	}
}
