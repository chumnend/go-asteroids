package asteroids

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	dpi             = 72
	fontSize        = 8
	menuBGColor     = color.RGBA{77, 170, 182, 200}
	textNormalColor = color.RGBA{34, 32, 32, 255}
)

type MenuState int

const (
	MenuMain MenuState = iota
	MenuPause
	MenuGameOver
	MenuWin
)

type textItem struct {
	text  string
	posY  int
	color color.Color
}

var mainMenuTexts = []textItem{
	{"Ebiten Asteroids", gameHeight / 2, textNormalColor},
	{"Press Space to start", gameHeight/2 + fontSize*2, textNormalColor},
}

var pauseTexts = []textItem{
	{"PAUSED", gameHeight / 2, textNormalColor},
	{"Press Space to continue", gameHeight/2 + fontSize*2, textNormalColor},
}

var gameOverTexts = []textItem{
	{"Game Over", gameHeight / 2, textNormalColor},
	{"Press Space to replay", gameHeight/2 + fontSize*2, textNormalColor},
}

var wonTexts = []textItem{
	{"You Won!", gameHeight / 2, textNormalColor},
	{"Press Space to replay", gameHeight/2 + fontSize*2, textNormalColor},
}

func (g *Game) loadMenuResources() error {
	// load the font type
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		return err
	}

	tf, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(fontSize),
		DPI:     float64(dpi),
		Hinting: font.HintingFull,
	})
	if err != nil {
		return err
	}

	g.font = tf

	return nil
}

func (g *Game) drawMenu(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, 0, 0, gameWidth, gameHeight, menuBGColor)

	// get the text items based on menu state
	var texts []textItem
	switch g.menuState {
	case MenuMain:
		texts = mainMenuTexts
	case MenuPause:
		texts = pauseTexts
	case MenuGameOver:
		texts = gameOverTexts
	case MenuWin:
		texts = wonTexts
	default:
		panic("unexpected state")
	}

	// draw each text item to the screen
	for _, ti := range texts {
		text.Draw(screen, ti.text, g.font, gameWidth/2.-len(ti.text)/2.*fontSize, ti.posY, ti.color)
	}
}
