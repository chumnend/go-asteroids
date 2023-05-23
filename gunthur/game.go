package gunthur

import (
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	playerImage     *ebiten.Image
	backgroundImage *ebiten.Image
)

const (
	screenWidth  = 320
	screenHeight = 256
	scale        = 2

	WindowWidth  = screenWidth * scale
	WindowHeight = screenHeight * scale
)

func init() {
	var err error

	backgroundImage, _, err = ebitenutil.NewImageFromFile("assets/backgrounds/bg_320x256.png")
	if err != nil {
		log.Fatal(err)
	}

	playerImage, _, err = ebitenutil.NewImageFromFile("assets/sprites/adventurer/adventurer-sheet.png")
	if err != nil {
		log.Fatal(err)
	}
}

type GameState int

const (
	GameStateMenu = iota
	GameStatePlaying
)

type Game struct {
	count int
	state GameState
}

func (g *Game) Update() error {
	g.count++

	switch g.state {
	case GameStateMenu:
		if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
			g.state = GameStatePlaying
		}
	case GameStatePlaying:
		if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
			g.state = GameStateMenu
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(backgroundImage, nil)
	drawAnimation(screen, playerImage, 50, 37, 50, 37, 6, g.count)
	// debug message
	ebitenutil.DebugPrint(screen, "Current state: "+strconv.Itoa(int(g.state)))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
