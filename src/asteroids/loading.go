package asteroids

import (
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// LOADING ========================================================================================

// loadObjects loads all required assets for the game
func (game *Game) loadObjects() error {
	return nil
}

func (game *Game) loadMenuResources() error {
	// load the font type
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		return err
	}

	tf, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(FONT_SIZE),
		DPI:     float64(DPI),
		Hinting: font.HintingFull,
	})
	if err != nil {
		return err
	}

	game.font = tf

	return nil
}
