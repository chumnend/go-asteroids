package gunthur

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Spritesheet struct {
	Image       *ebiten.Image
	FrameWidth  int
	FrameHeight int
}

func NewSpritesheet(path string, frameWidth int, frameHeight int) *Spritesheet {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return &Spritesheet{
		Image:       img,
		FrameWidth:  frameWidth,
		FrameHeight: frameHeight,
	}
}

func (s *Spritesheet) GetWidth() int {
	return s.Image.Bounds().Dx()
}

func (s *Spritesheet) GetHeight() int {
	return s.Image.Bounds().Dy()
}
