package asteroids

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// HANDLE INPUT ===================================================================================

// handleInput reads key inputs and performs actions
func (game *Game) handleInput() error {
	// get pressed keys
	game.pressedKeys = inpututil.AppendPressedKeys(game.pressedKeys[:0])

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
