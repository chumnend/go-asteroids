package asteroids

const (
	INITIAL_SHIP_X            = GAME_WIDTH / 2
	INITIAL_SHIP_Y            = GAME_HEIGHT / 2
	INITIAL_SHIP_VX           = 0
	INITIAL_SHIP_VY           = 0
	INITIAL_DIRECTION_DEGREES = 0
	TURN_RATE                 = 5 // in degrees
)

type Ship struct {
	Entity
}

func NewShip() (*Ship, error) {
	ship := &Ship{
		Entity: NewEntity(),
	}

	// load the player sprite
	var err error
	ship.Sprite, err = NewSpriteFromImagePath("src/assets/sprites/player.png")
	if err != nil {
		return nil, err
	}

	// set initial ship parameters
	ship.Initialize()
	ship.IsRotatable = true

	return ship, nil
}

func (ship *Ship) Rotate(clockwise bool) {
	sign := 1.
	if !clockwise {
		sign *= -1
	}
	ship.Direction += sign * float64(TURN_RATE)
}

func (ship *Ship) Initialize() {
	ship.Position.X = INITIAL_SHIP_X
	ship.Position.Y = INITIAL_SHIP_Y
	ship.Velocity.X = INITIAL_SHIP_VX
	ship.Velocity.Y = INITIAL_SHIP_VY
	ship.Direction = INITIAL_DIRECTION_DEGREES
}

func (ship *Ship) Update() {}
