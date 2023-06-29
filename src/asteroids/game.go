package asteroids

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type GameState int

const (
	GameStateMenu = iota
	GameStatePlaying
	GameStateGameOver
)

type Game struct {
	width  int
	height int

	state       GameState
	pressedKeys []ebiten.Key
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
	}
}

func (g *Game) Update() error {
	g.processInput()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case GameStateMenu:
		g.drawStartMenu(screen)
	case GameStatePlaying:
		t := "Score: 0"
		x := len(t) / 2
		text.Draw(screen, t, textFont, x, fontSize, color.White)
		purpleCol := color.RGBA{255, 0, 255, 255}
		for x := 125; x < 175; x++ {
			for y := 125; y < 175; y++ {
				screen.Set(x, y, purpleCol)
			}
		}
	case GameStateGameOver:
		g.drawGameOver(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}
