package main

import (
	"bytes"
	"image"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	rplatformer "github.com/hajimehoshi/ebiten/v2/examples/resources/images/platformer"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	// game settings
	SCREEN_WIDTH  = 320
	SCREEN_HEIGHT = 256
	SCALE         = 2
	WINDOW_WIDTH  = SCREEN_WIDTH * SCALE
	WINDOW_HEIGHT = SCREEN_HEIGHT * SCALE
)

var (
	leftSprite      *ebiten.Image
	rightSprite     *ebiten.Image
	idleSprite      *ebiten.Image
	backgroundImage *ebiten.Image
)

func init() {
	// preload images
	img, _, err := image.Decode(bytes.NewReader(rplatformer.Right_png))
	if err != nil {
		panic(err)
	}
	rightSprite = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(rplatformer.Left_png))
	if err != nil {
		panic(err)
	}
	leftSprite = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(rplatformer.MainChar_png))
	if err != nil {
		panic(err)
	}
	idleSprite = ebiten.NewImageFromImage(img)

	backgroundImage, _, err = ebitenutil.NewImageFromFile("assets/bg_320x256.png")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	g := &Game{}

	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Gophy Runner")
	ebiten.SetWindowResizable(false)

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

type GameState int

const (
	GameStateMenu = iota
	GameStatePlaying
)

type Game struct {
	state GameState
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
	// draw background
	screen.DrawImage(backgroundImage, nil)

	// draw gopher
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.2, 0.2)
	op.GeoM.Translate(SCREEN_WIDTH/2, SCREEN_WIDTH/2)
	screen.DrawImage(idleSprite, op)

	// debug message
	ebitenutil.DebugPrint(screen, "Current state: "+strconv.Itoa(int(g.state)))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
