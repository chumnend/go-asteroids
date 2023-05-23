package gunthur

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func drawAnimation(
	screen *ebiten.Image,
	texture *ebiten.Image,
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
	screen.DrawImage(texture.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}
