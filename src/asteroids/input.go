package asteroids

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) processInput() {
	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])

	switch g.state {
	case GameStateMenu:
		for _, key := range g.pressedKeys {
			switch key {
			case ebiten.KeySpace:
				g.state = GameStatePlaying
			}
		}
	case GameStatePlaying:
		for _, key := range g.pressedKeys {
			switch key {
			case ebiten.KeyP:
				g.state = GameStateGameOver
			}
		}
	case GameStateGameOver:
		for _, key := range g.pressedKeys {
			switch key {
			case ebiten.KeyR:
				g.state = GameStateMenu
			}
		}
	}
}
