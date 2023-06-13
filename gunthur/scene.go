package gunthur

import "github.com/hajimehoshi/ebiten/v2"

type Point struct {
	X, Y float64
}

// Scene implements a barebone container for components
type Scene struct {
	Components []interface{}
	Offset     Point
}

func (s *Scene) Update() error {
	// Find all the components that can be updated, and update them.
	for _, c := range s.Components {
		if u, ok := c.(Updater); ok {
			if err := u.Update(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Scene) Draw(screen *ebiten.Image, opts ebiten.DrawImageOptions) {
	// Copy opts
	newOpts := opts

	// Reset GeoM to an identity matrix
	newOpts.GeoM.Reset()

	// Translate from spritespace to worldspace
	newOpts.GeoM.Translate(s.Offset.X, s.Offset.Y)

	// Reapply opts.GeoM to translate from worldspace to screenspace via
	// whatever geometry matrix was supplied by the parent
	newOpts.GeoM.Concat(opts.GeoM)

	// Find all the components that can be drawn, and draw them.
	for _, c := range s.Components {
		if d, ok := c.(Drawer); ok {
			d.Draw(screen, newOpts)
		}
	}
}
