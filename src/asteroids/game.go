package asteroids

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GameState int

const (
	GameStateMenu = iota
	GameStatePlaying
	GameStateGameOver
)

// Game implements ebiten.Game interface
type Game struct {
	width  int
	height int

	state       GameState
	pressedKeys []ebiten.Key
}

// NewGame returns a Game struct, takes the size of the game screen
func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
	}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.processInput()
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case GameStateMenu:
		g.drawStartMenu(screen)
	case GameStatePlaying:
	case GameStateGameOver:
		g.drawGameOver(screen)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}
