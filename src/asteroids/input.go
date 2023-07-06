package asteroids

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) processInput() {
	switch g.gameState {
	case GameStateMenu:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.startGame()
		}
	case GameStatePlaying:
		if inpututil.IsKeyJustPressed(ebiten.KeyP) {
			g.pauseGame()
		}

		g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])

		for _, key := range g.pressedKeys {
			switch key.String() {
			case "ArrowDown":
				g.ship.Position.Y += 2
			case "ArrowUp":
				g.ship.Position.Y -= 2
			case "ArrowRight":
				g.ship.Position.X += 2
			case "ArrowLeft":
				g.ship.Position.X -= 2
			}
		}
	}
}
