package asteroids

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	fontSizeMd = 12
	fontSizeLg = fontSizeMd * 1.5
	dpi        = 72
)

var (
	tfMd font.Face
	tfLg font.Face
)

func init() {
	var err error

	// initialize fonts to be used
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}

	tfMd, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSizeMd,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	tfLg, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSizeLg,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}
