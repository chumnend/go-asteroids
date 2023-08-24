package asteroids

import (
	"errors"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	GAME_WIDTH    = 300 // internal game width
	GAME_HEIGHT   = 300 // internal game height
	SCALE         = 2   // factor to scale for desktop
	WINDOW_WIDTH  = GAME_WIDTH * SCALE
	WINDOW_HEIGHT = GAME_HEIGHT * SCALE
	DT            = 1 / 60.0 // assume the delta is fixed and we are always at 60 FPS
)

type GameState int

const (
	GameStateMenu GameState = iota
	GameStatePlaying
)

type MenuState int

const (
	MenuMain MenuState = iota
	MenuPause
	MenuGameOver
)

// Game implements the ebiten.Game interface
type Game struct {
	gameState GameState
	menuState MenuState
}

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

	// debug info
	game.printDebugInfo(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return GAME_WIDTH, GAME_HEIGHT
}

// LOADING ========================================================================================

// loadObjects loads all required assets for the game
func (game *Game) loadObjects() error {
	return nil
}

// HANDLE INPUT ===================================================================================

// handleInput reads key inputs and performs actions
func (game *Game) handleInput() error {
	// force game end
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	switch game.gameState {
	case GameStateMenu:
		switch game.menuState {
		case MenuMain:
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
				game.startGame()
			}
		case MenuPause:
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
				game.resumeGame()
			}
		default:
			return errors.New("unexpected menu state")
		}
	case GameStatePlaying:
		if inpututil.IsKeyJustPressed(ebiten.KeyP) {
			game.pauseGame()
		}
	default:
		return errors.New("unexpected game state")
	}

	return nil
}

func (game *Game) startGame() {
	game.gameState = GameStatePlaying
}

func (game *Game) pauseGame() {
	game.gameState = GameStateMenu
	game.menuState = MenuPause
}

func (game *Game) resumeGame() {
	game.gameState = GameStatePlaying
}

// PROCESS GAME LOGIC =============================================================================

// processLogic updates all game objects each frame
func (game *Game) processLogic() error {
	return nil
}

// PAINT SCREEN ===================================================================================

// drawObjects redraws all objects on the screen
func (game *Game) drawObjects() error {
	return nil
}

func (game *Game) printDebugInfo(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("State: %v\nMenu: %v", game.gameState, game.menuState))
}
