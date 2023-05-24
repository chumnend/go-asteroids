package gunthur

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	frameWidth    = 50
	frameHeight   = 37
	spirtesPerRow = 7
	sheetWidth    = frameWidth * spirtesPerRow

	idleFrameX     = 0
	idleFrameY     = 0
	idleFrameCount = 4

	runFrameX     = 50
	runFrameY     = 37
	runFrameCount = 6

	jumpFrameX     = 0
	jumpFrameY     = 74
	jumpFrameCount = 10

	crouchFrameX     = 200
	crouchFrameY     = 0
	crouchFrameCount = 4
)

type PlayerState int

const (
	PlayerStateIdle = iota
	PlayerStateRunRight
	PlayerStateRunLeft
	PlayerStateJumping
	PlayerStateCrouch
)

type Player struct {
	textureAtlas *ebiten.Image
	x            float64
	y            float64
	state        PlayerState
}

func NewPlayer(x float64, y float64) *Player {
	p := &Player{
		x: x,
		y: y,
	}

	var err error
	p.textureAtlas, _, err = ebitenutil.NewImageFromFile("assets/sprites/adventurer/adventurer-sheet.png")
	if err != nil {
		log.Fatal(err)
	}

	return p
}

func (p *Player) draw(screen *ebiten.Image, currentCount int) {
	switch p.state {
	case PlayerStateJumping:
		p.jump(screen, currentCount)
		break
	case PlayerStateCrouch:
		p.crouch(screen, currentCount)
		break
	case PlayerStateRunRight:
		p.runRight(screen, currentCount)
		break
	case PlayerStateRunLeft:
		p.runLeft(screen, currentCount)
		break
	default:
		p.idle(screen, currentCount)
		break
	}
}

func (p *Player) idle(screen *ebiten.Image, currentCount int) {
	p.animate(screen, currentCount, idleFrameX, idleFrameY, idleFrameCount, false)
}

func (p *Player) runRight(screen *ebiten.Image, currentCount int) {
	p.animate(screen, currentCount, runFrameX, runFrameY, runFrameCount, false)
}

func (p *Player) runLeft(screen *ebiten.Image, currentCount int) {
	p.animate(screen, currentCount, runFrameX, runFrameY, runFrameCount, true)
}

func (p *Player) jump(screen *ebiten.Image, currentCount int) {
	p.animate(screen, currentCount, jumpFrameX, jumpFrameY, jumpFrameCount, false)
}

func (p *Player) crouch(screen *ebiten.Image, currentCount int) {
	p.animate(screen, currentCount, crouchFrameX, crouchFrameY, crouchFrameCount, false)
}

func (p *Player) animate(screen *ebiten.Image, currentCount int, startFrameX int, startFrameY int, frameCount int, mirror bool) {
	i := (currentCount / frameCount) % frameCount
	sx := startFrameX + i*frameWidth
	sy := startFrameY
	if sx >= sheetWidth {
		timesOverflow := int(sx / sheetWidth)
		sx = sx - sheetWidth*timesOverflow
		sy = sy + frameHeight*timesOverflow
	}

	img := p.textureAtlas.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.x, p.y)
	if mirror {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(screenWidth+frameWidth, 0)
	}

	screen.DrawImage(img, op)
}
