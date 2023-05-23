package main

import (
	"gophy-runner/gunthur"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := &gunthur.Game{}

	ebiten.SetWindowSize(gunthur.WINDOW_WIDTH, gunthur.WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Gunthur: Hack & Slash")
	ebiten.SetWindowResizable(false)

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
