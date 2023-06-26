package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	// internal game size (before scaling)
	gameScreenWidth  = 320
	gameScreenHeight = 240

	scale        = 2 // scale 200% in window
	windowWidth  = gameScreenWidth * scale
	windowHeight = gameScreenHeight * scale

	fontSize    = 12
	fontSizeLrg = fontSize * 1.5
)

var (
	titleFont font.Face
	textFont  font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	titleFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSizeLrg,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	textFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

type GameState int

const (
	GameStateMenu = iota
	GameStatePlaying
	GameStateGameOver
)

type Game struct {
	state GameState
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var titleTexts []string
	var texts []string

	switch g.state {
	case GameStateMenu:
		menuBgColor := color.RGBA{0, 0, 0, 255}
		ebitenutil.DrawRect(screen, 0, 0, gameScreenWidth, gameScreenHeight, menuBgColor)
		titleTexts = []string{"ASTEROIDS"}
		texts = []string{"Press Space to Start"}

		for i, t := range titleTexts {
			x := (gameScreenWidth - len(t)*fontSizeLrg) / 2
			text.Draw(screen, t, titleFont, x, (i+4)*fontSizeLrg, color.White)
		}
		for i, t := range texts {
			x := (gameScreenWidth - len(t)*fontSize) / 2
			text.Draw(screen, t, textFont, x, (i+4)*2*fontSize, color.White)
		}
	case GameStatePlaying:
	case GameStateGameOver:
		texts = []string{"", "GAME OVER!"}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return gameScreenWidth, gameScreenHeight
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Gunthur: High-Speed Hack and Slash")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
