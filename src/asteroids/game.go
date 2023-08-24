package asteroids

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	GAME_WIDTH    = 300 // internal game width
	GAME_HEIGHT   = 300 // internal game height
	SCALE         = 2   // factor to scale for desktop
	WINDOW_WIDTH  = GAME_WIDTH * SCALE
	WINDOW_HEIGHT = GAME_HEIGHT * SCALE
	DT            = 1 / 60.0 // assume the delta is fixed and we are always at 60 FPS
)

// Game implements the ebiten.Game interface
type Game struct{}

// NewGame returns a Game struct, the width of the window and the height of the window
func NewGame() (*Game, int, int) {
	return &Game{}, WINDOW_WIDTH, WINDOW_HEIGHT
}

// Init loads all resources for the game
func (game *Game) Init() error {
	if err := game.loadObjects(); err != nil {
		return err
	}

	return nil
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (game *Game) Update() error {
	if err := game.handleInput(); err != nil {
		return err
	}

	if err := game.processLogic(); err != nil {
		return err
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (game *Game) Draw(screen *ebiten.Image) {
	if err := game.drawObjects(); err != nil {
		log.Fatal(err)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return GAME_WIDTH, GAME_HEIGHT
}
