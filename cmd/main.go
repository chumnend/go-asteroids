package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	SCREEN_WIDTH  = 300
	SCREEN_HEIGHT = 300
)

type Game struct{}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "This is a test.")
}

func main() {
	ebiten.SetWindowSize(SCREEN_WIDTH*2, SCREEN_HEIGHT*2)
	ebiten.SetWindowTitle("Ebiten Test")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
