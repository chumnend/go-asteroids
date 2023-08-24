package asteroids

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// PROCESS GAME LOGIC =============================================================================

// processLogic updates all game objects each frame
func (game *Game) processLogic() error {
	return nil
}

// PAINT SCREEN ===================================================================================

func (g *Game) drawMenuScreen(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, 0, 0, GAME_WIDTH, GAME_HEIGHT, MENU_BG_COLOR)

	// get the text items based on menu state
	var texts []textItem
	switch g.menuState {
	case MenuMain:
		texts = mainMenuTexts
	case MenuPause:
		texts = pauseTexts
	case MenuGameOver:
		texts = gameOverTexts
	case MenuWin:
		texts = wonTexts
	default:
		panic("unexpected state")
	}

	// draw each text item to the screen
	for _, ti := range texts {
		text.Draw(screen, ti.text, g.font, GAME_WIDTH/2.-len(ti.text)/2.*FONT_SIZE, ti.posY, ti.color)
	}
}

func (game *Game) printDebugInfo(screen *ebiten.Image) {
	var currentGameState string
	switch game.gameState {
	case GameStateMenu:
		currentGameState = "Menu"
	case GameStatePlaying:
		currentGameState = "Playing"
	}

	var currentMenuState string
	switch game.menuState {
	case MenuMain:
		currentMenuState = "Main Menu"
	case MenuPause:
		currentMenuState = "Paused"
	case MenuGameOver:
		currentMenuState = "Game Over"
	case MenuWin:
		currentMenuState = "Win"
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("State: %v   Menu: %v\nKeys: %v", currentGameState, currentMenuState, game.pressedKeys))
}
