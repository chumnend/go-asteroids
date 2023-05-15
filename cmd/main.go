package main

import (
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	SCREEN_WIDTH  = 320
	SCREEN_HEIGHT = 256

	SCALE = 2

	WINDOW_WIDTH  = SCREEN_WIDTH * SCALE
	WINDOW_HEIGHT = SCREEN_HEIGHT * SCALE
)

type GameState int

const (
	GameStateMenu = iota
	GameStatePlaying
)

type Game struct {
	bg *ebiten.Image

	state GameState
}

func (g *Game) Init() error {
	var err error
	g.bg, _, err = ebitenutil.NewImageFromFile("assets/bg_320x256.png")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (g *Game) Update() error {
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
	screen.DrawImage(g.bg, nil)
	ebitenutil.DebugPrint(screen, "Current state: "+strconv.Itoa(int(g.state)))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	g := &Game{}
	if err := g.Init(); err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Gophy Runner")
	ebiten.SetWindowResizable(false)

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
