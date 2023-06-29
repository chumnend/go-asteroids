package main

import (
	"gophy-runner/src/asteroids"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// internal game size (before scaling)
	gameScreenWidth  = 500
	gameScreenHeight = 300

	scale        = 2 // scale 200% in window
	windowWidth  = gameScreenWidth * scale
	windowHeight = gameScreenHeight * scale
)

func main() {
	game := asteroids.NewGame(gameScreenWidth, gameScreenHeight)

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Asteroids")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
