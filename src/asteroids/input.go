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

		rect := g.ship.getAABB()

		for _, key := range g.pressedKeys {
			switch key.String() {
			case "ArrowUp":
				newPos := g.ship.Position.Y - 2
				if newPos > 0 {
					g.ship.Position.Y = newPos
				}
			case "ArrowDown":
				newPos := g.ship.Position.Y + 2
				if newPos+int(rect.H) < gameHeight {
					g.ship.Position.Y = newPos
				}
			case "ArrowLeft":
				newPos := g.ship.Position.X - 2
				if newPos > 0 {
					g.ship.Position.X = newPos
				}
			case "ArrowRight":
				newPos := g.ship.Position.X + 2
				if newPos+int(rect.W) < gameWidth {
					g.ship.Position.X = newPos
				}
			}
		}
	}
}
