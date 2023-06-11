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
	Components  []interface{}
	pressedKeys []ebiten.Key
}

type IDraw interface {
	Draw(*ebiten.Image)
}

type IUpdate interface {
	Update() error
}

type IHandleInput interface {
	HandleInput(key []ebiten.Key)
}

// NewVector returns a Vector struct, window width and window height
func NewGame() (*Game, int, int) {
	g := &Game{}

	player := NewSprite()
	g.Components = append(g.Components, player)

	return g, screenWidth * scale, screenHeight * scale
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])
	for _, c := range g.Components {
		if h, ok := c.(IHandleInput); ok {
			h.HandleInput(g.pressedKeys)
		}
	}

	// Find all the components that can be updated, and update them.
	for _, c := range g.Components {
		if u, ok := c.(IUpdate); ok {
			if err := u.Update(); err != nil {
				return err
			}
		}
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Find all the components that can be drawn, and draw them.
	for _, c := range g.Components {
		if d, ok := c.(IDraw); ok {
			d.Draw(screen)
		}
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
