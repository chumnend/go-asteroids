package gunthur

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Player represents a game sprite
type Player struct {
	x      int
	y      int
	width  int
	height int
}

// NewPlayer returns a Player struct
func NewPlayer() *Player {
	return &Player{
		x:      200,
		y:      200,
		width:  50,
		height: 50,
	}
}

// MoveTo changes the position of the sprite
func (p *Player) MoveTo(x int, y int) {
	p.x = x
	if p.x < 0 {
		p.x = 0
	} else if x > screenWidth-p.width {
		p.x = screenWidth - p.width
	}

	p.y = y
	if p.y < 0 {
		p.y = 0
	} else if p.y > screenHeight-p.height {
		p.y = screenHeight - p.height
	}
}

func (p *Player) HandleInput(keys []ebiten.Key) {
	for _, key := range keys {
		switch key.String() {
		case "ArrowDown":
			p.MoveTo(p.x, p.y+5)
		case "ArrowUp":
			p.MoveTo(p.x, p.y-5)
		case "ArrowRight":
			p.MoveTo(p.x+5, p.y)
		case "ArrowLeft":
			p.MoveTo(p.x-5, p.y)
		}
	}
}

func (p *Player) Update() error {
	return nil
}

// Draw takes an ebiten screen object and draws the sprite on it
func (p *Player) Draw(screen *ebiten.Image, opts ebiten.DrawImageOptions) {
	ebitenutil.DebugPrint(screen, "X: "+strconv.Itoa(p.x)+" Y:"+strconv.Itoa(p.y))

	purple := color.RGBA{255, 0, 255, 255}
	for x := p.x; x < p.x+p.width; x++ {
		for y := p.y; y < p.y+p.height; y++ {
			screen.Set(x, y, purple)
		}
	}
}
