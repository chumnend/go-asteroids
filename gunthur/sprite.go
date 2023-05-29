package gunthur

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Sprite manages info related to a game image
type Sprite struct {
	currentAnimation string

	Spritesheet Spritesheet
	Animations  map[string]*Animation
	X           int
	Y           int
}

// New Sprite returns a Sprite struct
func NewSprite(spritesheet *Spritesheet) *Sprite {
	return &Sprite{
		currentAnimation: "default",
		Spritesheet:      *spritesheet,
		Animations:       make(map[string]*Animation),
		X:                100,
		Y:                100,
	}
}

func (s *Sprite) AddAnimation(name string, duration int, startFrameX int, startFrameY int, totalFrames int, mirrorX bool, mirrorY bool) {
	animation := NewAnimation(duration, startFrameX, startFrameY, totalFrames, mirrorX, mirrorY)
	s.Animations[name] = animation
}

func (s *Sprite) SetAnimation(name string) {
	s.currentAnimation = name
}

func (s *Sprite) Draw(screen *ebiten.Image, count int) {
	currentAnimation := s.Animations[s.currentAnimation]

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.X), float64(s.Y))
	if currentAnimation.MirrorY {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(s.Spritesheet.GetWidth()-s.X), 0)
	}

	i := (count / currentAnimation.TotalFrames) % currentAnimation.TotalFrames
	sx := currentAnimation.StartFrameX + i*s.Spritesheet.FrameWidth
	sy := currentAnimation.StartFrameY
	if sx >= s.Spritesheet.GetWidth() {
		timesOverflow := int(sx / s.Spritesheet.GetWidth())
		sx = sx - s.Spritesheet.GetWidth()*timesOverflow
		sy = sy + s.Spritesheet.FrameHeight*timesOverflow
	}

	img := s.Spritesheet.Image.SubImage(image.Rect(sx, sy, sx+s.Spritesheet.FrameWidth, sy+s.Spritesheet.FrameHeight)).(*ebiten.Image)

	screen.DrawImage(img, op)
}
