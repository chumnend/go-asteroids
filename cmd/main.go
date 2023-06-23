package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	// internal game size (before scaling)
	gameScreenWidth  = 320
	gameScreenHeight = 240

	scale        = 2 // scale 200% in window
	windowWidth  = gameScreenWidth * scale
	windowHeight = gameScreenHeight * scale
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return gameScreenWidth, gameScreenHeight
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Gunthur: High-Speed Hack and Slash")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
