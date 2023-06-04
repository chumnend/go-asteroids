package gunthur

import (
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Sprite manages info related to a game image
type Sprite struct {
	currentAnimation string
	startTime        time.Time

	Spritesheet Spritesheet
	Animations  map[string]*Animation
	X           int
	Y           int
	Vx          int
	Vy          int
}

// New Sprite returns a Sprite struct
func NewSprite(spritesheet *Spritesheet) *Sprite {
	return &Sprite{
		currentAnimation: "default",
		Spritesheet:      *spritesheet,
		Animations:       make(map[string]*Animation),
		X:                100,
		Y:                100,
		Vx:               0,
		Vy:               0,
	}
}

func (s *Sprite) AddAnimation(name string, duration int, startFrameX int, startFrameY int, totalFrames int, mirrorX bool, mirrorY bool) {
	animation := NewAnimation(duration, startFrameX, startFrameY, totalFrames, mirrorX, mirrorY)
	s.Animations[name] = animation
}

func (s *Sprite) SetAnimation(name string) {
	if s.currentAnimation != name {
		s.currentAnimation = name
		s.startTime = time.Now()
	}
}

func (s *Sprite) MoveTo(x int, y int) {
	s.X = x
	s.Y = y
}

func (s *Sprite) SetSpeed(vx int, vy int) {
	s.Vx = vx
	s.Vy = vy
}

func (s *Sprite) UpdatePosition() {
	s.X += s.Vx
	s.Y += s.Vy
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	currentAnimation := s.Animations[s.currentAnimation]
	elapsedTime := float32(time.Since(s.startTime).Milliseconds())
	oneStepDuration := currentAnimation.Duration.Milliseconds() / int64(currentAnimation.TotalFrames)
	i := int(elapsedTime) / int(oneStepDuration)
	if elapsedTime >= float32(currentAnimation.Duration.Milliseconds()) {
		s.startTime = time.Now()
	}

	sx := currentAnimation.StartFrameX + i*s.Spritesheet.FrameWidth
	sy := currentAnimation.StartFrameY
	if sx >= s.Spritesheet.GetWidth() {
		timesOverflow := int(sx / s.Spritesheet.GetWidth())
		sx = sx - s.Spritesheet.GetWidth()*timesOverflow
		sy = sy + s.Spritesheet.FrameHeight*timesOverflow
	}

	img := s.Spritesheet.Image.SubImage(image.Rect(sx, sy, sx+s.Spritesheet.FrameWidth, sy+s.Spritesheet.FrameHeight)).(*ebiten.Image)

	op := &ebiten.DrawImageOptions{}
	scaleX := 1.
	scaleY := 1.
	offsetX := float64(s.X)
	offsetY := float64(s.Y)
	if currentAnimation.MirrorY {
		scaleX = -1
		offsetX += float64(s.Spritesheet.FrameWidth)
	}
	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(offsetX, offsetY)

	screen.DrawImage(img, op)
}
