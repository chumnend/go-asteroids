package asteroids

import "github.com/hajimehoshi/ebiten/v2"

// Scene implements a barebone container for components
type Scene struct {
	components []interface{}
	offset     Position
}

// NewScene returns a Scene struct, with an offset from parent
func NewScene() *Scene {
	return &Scene{}
}

func (s *Scene) AddComponent(component interface{}) {
	s.components = append(s.components, component)
}

func (s *Scene) SetOffset(x, y int) {
	s.offset.x = x
	s.offset.y = y
}

func (s *Scene) Update(keys []ebiten.Key) error {
	// Find all the components that can handle input and pass them the keys.
	for _, c := range s.components {
		if h, ok := c.(Inputter); ok {
			h.HandleInput(keys)
		}
	}

	// Find all the components that can be updated, and update them.
	for _, c := range s.components {
		if u, ok := c.(Updater); ok {
			if err := u.Update(keys); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Scene) Draw(screen *ebiten.Image, opts ebiten.DrawImageOptions) {
	// Take parent opts and apply scene offsets
	newOpts := opts                                                  // Copy opts
	newOpts.GeoM.Reset()                                             // Reset GeoM to an identity matrix
	newOpts.GeoM.Translate(float64(s.offset.x), float64(s.offset.y)) // Translate from spritespace to worldspace
	newOpts.GeoM.Concat(opts.GeoM)                                   // Reapply opts.GeoM to translate from worldspace to screenspace via whatever geometry matrix was supplied by the parent

	// Find all the components that can be drawn, and draw them.
	for _, c := range s.components {
		if d, ok := c.(Drawer); ok {
			d.Draw(screen, newOpts)
		}
	}
}
