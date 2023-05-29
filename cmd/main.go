package main

import (
	"gophy-runner/gunthur"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := gunthur.NewGame()
	ebiten.SetWindowSize(gunthur.WindowWidth, gunthur.WindowHeight)
	ebiten.SetWindowTitle("Gunthur: Hack & Slash")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
