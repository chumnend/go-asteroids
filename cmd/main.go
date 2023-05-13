package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	SCREEN_HEIGHT = 480
	SCREEN_WIDTH  = 640
)

var (
	err        error
	background *ebiten.Image
	spaceShip  *ebiten.Image
	player     Player
)

type Player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

func init() {
	background, _, err = ebitenutil.NewImageFromFile("assets/space-background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	spaceShip, _, err = ebitenutil.NewImageFromFile("assets/spaceship.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	player = Player{
		image: spaceShip,
		xPos:  SCREEN_WIDTH / 2,
		yPos:  SCREEN_HEIGHT / 2,
		speed: 4,
	}
}

func movePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		player.yPos -= player.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		player.yPos += player.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		player.xPos -= player.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		player.xPos += player.speed
	}
}

func update(screen *ebiten.Image) error {
	movePlayer()
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(0, 0)
	screen.DrawImage(background, opt)

	playerOpt := &ebiten.DrawImageOptions{}
	playerOpt.GeoM.Translate(player.xPos, player.yPos)
	screen.DrawImage(player.image, playerOpt)

	return nil
}

func main() {
	if err := ebiten.Run(update, SCREEN_WIDTH, SCREEN_HEIGHT, 1, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
