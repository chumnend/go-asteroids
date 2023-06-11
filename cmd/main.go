package main

import (
	"gophy-runner/gunthur"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, w, h := gunthur.NewGame()
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("Gunthur: The Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
