package main

import (
	"gophy-runner/src/asteroids"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g, w, h := asteroids.NewGame()
	if err := g.Init(); err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("Asteroids - Made With Go")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

	if err := g.Exit(); err != nil {
		log.Fatal(err)
	}
}
