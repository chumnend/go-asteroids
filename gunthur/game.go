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

// Inputter is an interface for components that are to receive user input
type Inputter interface {
	HandleInput([]ebiten.Key)
}

// Updater is an interface for components that are to be updated
type Updater interface {
	Update([]ebiten.Key) error
}

// Drawer is an interface for components that are to be drawn to ebiten screen
type Drawer interface {
	Draw(*ebiten.Image, ebiten.DrawImageOptions)
}

// Game implements ebiten.Game interface
type Game struct {
	Scene

	pressedKeys []ebiten.Key
}

// NewVector returns a Vector struct, window width and window height
func NewGame() (*Game, int, int) {
	g := &Game{}
	g.Scene.Components = append(g.Scene.Components, NewSprite()) // add player sprite to game

	return g, screenWidth * scale, screenHeight * scale
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])
	g.Scene.Update(g.pressedKeys)

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.Scene.Draw(screen, ebiten.DrawImageOptions{})
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
