package gunthur

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	textureAtlas *ebiten.Image
}

func NewPlayer() *Player {
	p := &Player{}

	var err error
	p.textureAtlas, _, err = ebitenutil.NewImageFromFile("assets/sprites/adventurer/adventurer-sheet.png")
	if err != nil {
		log.Fatal(err)
	}

	return p
}

func (p *Player) drawRunAninmation(
	screen *ebiten.Image,
	startFrameX int,
	startFrameY int,
	frameWidth int,
	frameHeight int,
	frameCount int,
	currentCount int,
) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	i := (currentCount / 5) % frameCount
	sx, sy := startFrameX+i*frameWidth, startFrameY
	screen.DrawImage(p.textureAtlas.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}
