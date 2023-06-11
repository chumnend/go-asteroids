package gunthur

import "github.com/hajimehoshi/ebiten/v2"

// Scene implements a barebone container for components
type Scene struct {
	Components []interface{}
}

func (s *Scene) Draw(screen *ebiten.Image) {
	for _, c := range s.Components {
		if d, ok := c.(IDraw); ok {
			d.Draw(screen)
		}
	}
}

func (s *Scene) Update() error {
	for _, c := range s.Components {
		if u, ok := c.(IUpdate); ok {
			if err := u.Update(); err != nil {
				return err
			}
		}
	}
	return nil
}
