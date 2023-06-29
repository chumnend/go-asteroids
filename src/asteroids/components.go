package asteroids

import "github.com/hajimehoshi/ebiten/v2"

type Position struct {
	x, y int
}

// Inputter is an interface for components that are to receive user input
type Inputter interface {
	HandleInput(keys []ebiten.Key)
}

// Updater is an interface for components that are to be updated
type Updater interface {
	Update(keys []ebiten.Key) error
}

// Drawer is an interface for components that are to be drawn to ebiten screen
type Drawer interface {
	Draw(screen *ebiten.Image, opts ebiten.DrawImageOptions)
}
