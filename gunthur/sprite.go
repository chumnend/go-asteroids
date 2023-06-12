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

func (s *Sprite) HandleInput(keys []ebiten.Key) {
	for _, key := range keys {
		switch key.String() {
		case "ArrowDown":
			s.MoveTo(s.x, s.y+5)
		case "ArrowUp":
			s.MoveTo(s.x, s.y-5)
		case "ArrowRight":
			s.MoveTo(s.x+5, s.y)
		case "ArrowLeft":
			s.MoveTo(s.x-5, s.y)
		}
	}
}

func (s *Sprite) Update() error {
	return nil
}

// Draw takes an ebiten screen object and draws the sprite on it
func (s *Sprite) Draw(screen *ebiten.Image, opts ebiten.DrawImageOptions) {
	purple := color.RGBA{255, 0, 255, 255}
	for x := s.x; x < s.x+s.width; x++ {
		for y := s.y; y < s.y+s.height; y++ {
			screen.Set(x, y, purple)
		}
	}
}
