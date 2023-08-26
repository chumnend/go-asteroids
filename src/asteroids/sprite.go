package asteroids

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	texture *ebiten.Image // fill spritesheet
	image   *ebiten.Image // portion of spritesheet to show
	mirrorX bool
	mirrorY bool
}

func NewSprite(img *ebiten.Image) *Sprite {
	return &Sprite{
		texture: img,
		image:   img,
	}
}

func NewSpriteFromImagePath(path string) (*Sprite, error) {
	// load sprite from file path
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return nil, err
	}
	sprite := NewSprite(img)

	return sprite, nil
}

func (s *Sprite) SetTextureRect(rect image.Rectangle) {
	s.image = s.texture.SubImage(rect).(*ebiten.Image)
}

func (s *Sprite) GetTextureRect() image.Rectangle {
	return s.image.Bounds()
}

func (s *Sprite) GetSize() Vector2 {
	return Vector2{
		X: float64(s.image.Bounds().Dx()), // width
		Y: float64(s.image.Bounds().Dy()), // height
	}
}

func (s *Sprite) Draw(screen *ebiten.Image, parentM ebiten.GeoM) {
	scale := Vector2{1., 1.}
	offset := Vector2{}

	imageSize := s.GetSize()

	if s.mirrorX {
		scale.X = -1
		offset.X = float64(imageSize.X)
	}
	if s.mirrorY {
		scale.Y = -1
		offset.Y = float64(imageSize.Y)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale.X, scale.Y)
	op.GeoM.Translate(offset.X, offset.Y)
	op.GeoM.Concat(parentM)

	screen.DrawImage(s.image, op)
}
