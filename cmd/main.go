package main

import (
	"gophy-runner/internal/engine"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	SCREEN_HEIGHT = 480
	SCREEN_WIDTH  = 640
)

func main() {
	game := &engine.Game{}
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("GOPHY RUNNER")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
