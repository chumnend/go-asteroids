package asteroids

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) processInput() {
	switch g.gameState {
	case GameStateMenu:
		switch g.menuState {
		case MenuMain:
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
				g.startGame()
			}
		case MenuPause:
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
				g.startGame()
			}
		case MenuGameOver:
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
				g.restartGame()
			}
		case MenuWin:
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
				g.restartGame()
			}
		default:
			panic("unexpected state")
		}

	case GameStatePlaying:
		if inpututil.IsKeyJustPressed(ebiten.KeyP) {
			g.pauseGame()
		}

		g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])

		for _, key := range g.pressedKeys {
			switch key {
			case ebiten.KeyArrowUp:
				g.ship.moveUp()
			case ebiten.KeyArrowDown:
				g.ship.moveDown()
			case ebiten.KeyArrowLeft:
				g.ship.moveLeft()
			case ebiten.KeyArrowRight:
				g.ship.moveRight()
			}
		}
	}
}
