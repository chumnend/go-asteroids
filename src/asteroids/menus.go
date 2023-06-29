package asteroids

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) drawStartMenu(screen *ebiten.Image) {
	var titleTexts []string
	var texts []string

	menuBgColor := color.RGBA{0, 0, 0, 255}
	ebitenutil.DrawRect(screen, 0, 0, float64(g.width), float64(g.height), menuBgColor)
	titleTexts = []string{"ASTEROIDS"}
	texts = []string{"Press Space to start"}

	for i, t := range titleTexts {
		x := (int(g.width) - len(t)*fontSizeLg) / 2
		text.Draw(screen, t, tfLg, x, (i+4)*fontSizeLg, color.White)
	}
	for i, t := range texts {
		x := (g.width - len(t)*fontSizeMd) / 2
		text.Draw(screen, t, tfMd, x, (i+4)*2*fontSizeMd, color.White)
	}
}

func (g *Game) drawGameOver(screen *ebiten.Image) {
	var titleTexts []string
	var texts []string

	menuBgColor := color.RGBA{0, 0, 0, 255}
	ebitenutil.DrawRect(screen, 0, 0, float64(g.width), float64(g.height), menuBgColor)
	titleTexts = []string{"Game Over"}
	texts = []string{"Press R to restart"}

	for i, t := range titleTexts {
		x := (g.width - len(t)*fontSizeLg) / 2
		text.Draw(screen, t, tfLg, x, (i+4)*fontSizeLg, color.White)
	}
	for i, t := range texts {
		x := (g.width - len(t)*fontSizeMd) / 2
		text.Draw(screen, t, tfMd, x, (i+4)*2*fontSizeMd, color.White)
	}
}