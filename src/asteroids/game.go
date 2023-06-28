package asteroids

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
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
	width  int
	height int

	state       GameState
	pressedKeys []ebiten.Key
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
	}
}

func (g *Game) processInput() {
	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])

	switch g.state {
	case GameStateMenu:
		for _, key := range g.pressedKeys {
			switch key {
			case ebiten.KeySpace:
				g.state = GameStatePlaying
			}
		}
	case GameStatePlaying:
		for _, key := range g.pressedKeys {
			switch key {
			case ebiten.KeyP:
				g.state = GameStateGameOver
			}
		}
	case GameStateGameOver:
		for _, key := range g.pressedKeys {
			switch key {
			case ebiten.KeyR:
				g.state = GameStateMenu
			}
		}
	}

}

func (g *Game) Update() error {
	g.processInput()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var titleTexts []string
	var texts []string

	switch g.state {
	case GameStateMenu:
		menuBgColor := color.RGBA{0, 0, 0, 255}
		ebitenutil.DrawRect(screen, 0, 0, float64(g.width), float64(g.height), menuBgColor)
		titleTexts = []string{"ASTEROIDS"}
		texts = []string{"Press Space to start"}

		for i, t := range titleTexts {
			x := (int(g.width) - len(t)*fontSizeLrg) / 2
			text.Draw(screen, t, titleFont, x, (i+4)*fontSizeLrg, color.White)
		}
		for i, t := range texts {
			x := (g.width - len(t)*fontSize) / 2
			text.Draw(screen, t, textFont, x, (i+4)*2*fontSize, color.White)
		}
	case GameStatePlaying:
		t := "Score: 0"
		x := len(t) / 2
		text.Draw(screen, t, textFont, x, fontSize, color.White)
		purpleCol := color.RGBA{255, 0, 255, 255}
		for x := 125; x < 175; x++ {
			for y := 125; y < 175; y++ {
				screen.Set(x, y, purpleCol)
			}
		}
	case GameStateGameOver:
		menuBgColor := color.RGBA{0, 0, 0, 255}
		ebitenutil.DrawRect(screen, 0, 0, float64(g.width), float64(g.height), menuBgColor)
		titleTexts = []string{"Game Over"}
		texts = []string{"Press R to restart"}

		for i, t := range titleTexts {
			x := (g.width - len(t)*fontSizeLrg) / 2
			text.Draw(screen, t, titleFont, x, (i+4)*fontSizeLrg, color.White)
		}
		for i, t := range texts {
			x := (g.width - len(t)*fontSize) / 2
			text.Draw(screen, t, textFont, x, (i+4)*2*fontSize, color.White)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}
