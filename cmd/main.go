package main

import (
	"gophy-runner/gunthur"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

func main() {
	g := gunthur.NewGame(screenWidth, screenHeight)
	if err := g.Init(); err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Gunthur: Hack & Slash")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
