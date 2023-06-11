package gunthur

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 450
	screenHeight = 400
	scale        = 2
)

// Game implements ebiten.Game interface
type Game struct {
	player      *Sprite
	pressedKeys []ebiten.Key
}

// NewVector returns a Vector struct, window width and window height
func NewGame() (*Game, int, int) {
	g := &Game{}
	g.player = NewSprite()

	return g, screenWidth * 2, screenHeight * 2
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])

	for _, key := range g.pressedKeys {
		switch key.String() {
		case "ArrowDown":
			g.player.MoveTo(g.player.x, g.player.y+5)
		case "ArrowUp":
			g.player.MoveTo(g.player.x, g.player.y-5)
		case "ArrowRight":
			g.player.MoveTo(g.player.x+5, g.player.y)
		case "ArrowLeft":
			g.player.MoveTo(g.player.x-5, g.player.y)
		}
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
