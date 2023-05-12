package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	err        error
	background *ebiten.Image
)

func init() {
	background, _, err = ebitenutil.NewImageFromFile("assets/space-background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(0, 0)
	screen.DrawImage(background, opt)

	return nil
}

func main() {
	if err := ebiten.Run(update, 640, 480, 1, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
