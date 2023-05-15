package engine

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

var (
	err          error
	background   *ebiten.Image
	spaceShip    *ebiten.Image
	player       Player
	screenHeight float64
	screenWidth  float64
)

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
		xPos:  screenWidth / 2,
		yPos:  screenHeight / 2,
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

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	movePlayer()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(0, 0)
	screen.DrawImage(background, opt)

	playerOpt := &ebiten.DrawImageOptions{}
	playerOpt.GeoM.Translate(player.xPos, player.yPos)
	screen.DrawImage(player.image, playerOpt)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	screenWidth = outsideWidth
	screenHeight = outsideHeight
	return screenWidth, screenHeight
}
