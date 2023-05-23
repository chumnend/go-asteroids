package gunthur

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	x  int
	y  int
	vx int
	vy int
}

func (p *Player) Draw(screen *ebiten.Image) {
	s := idleSprite
	switch {
	case p.vx > 0:
		s = rightSprite
	case p.vx < 0:
		s = leftSprite
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.2, 0.2)
	op.GeoM.Translate(float64(p.x)/SPRITE_SIZE, float64(p.y)/SPRITE_SIZE)
	screen.DrawImage(s, op)
}
