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
)

type Player struct {
	textureAtlas *ebiten.Image
	x            float64
	y            float64
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

func (p *Player) animate(screen *ebiten.Image, currentCount int, startFrameX int, startFrameY int, frameCount int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.x, p.y)

	i := (currentCount / frameCount) % frameCount
	sx := startFrameX + i*frameWidth
	sy := startFrameY
	if sx >= sheetWidth {
		timesOverflow := int(sx / sheetWidth)
		sx = sx - sheetWidth*timesOverflow
		sy = sy + frameHeight*timesOverflow
	}

	screen.DrawImage(p.textureAtlas.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func (p *Player) idle(screen *ebiten.Image, currentCount int) {
	p.animate(screen, currentCount, idleFrameX, idleFrameY, idleFrameCount)
}

func (p *Player) runRight(screen *ebiten.Image, currentCount int) {
	p.animate(screen, currentCount, runFrameX, runFrameY, runFrameCount)
}

func (p *Player) jump(screen *ebiten.Image, currentCount int) {
	p.animate(screen, currentCount, jumpFrameX, jumpFrameY, jumpFrameCount)
}
