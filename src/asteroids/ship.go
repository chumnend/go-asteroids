package asteroids

const (
	INITIAL_SHIP_X            = GAME_WIDTH / 2
	INITIAL_SHIP_Y            = GAME_HEIGHT / 2
	INITIAL_SHIP_VX           = 0
	INITIAL_SHIP_VY           = 0
	INITIAL_DIRECTION_DEGREES = 0
	TURN_RATE                 = 3 // represents angle in degrees
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
	ship.sprite, err = NewSpriteFromImagePath("src/assets/sprites/player.png")
	if err != nil {
		return nil, err
	}

	// set initial ship parameters
	ship.Initialize()
	ship.isRotatable = true

	return ship, nil
}

func (ship *Ship) rotate(clockwise bool) {
	sign := 1.
	if !clockwise {
		sign *= -1
	}
	ship.direction += sign * float64(TURN_RATE)
}

func (ship *Ship) Initialize() {
	ship.position.X = INITIAL_SHIP_X
	ship.position.Y = INITIAL_SHIP_Y
	ship.velocity.X = INITIAL_SHIP_VX
	ship.velocity.Y = INITIAL_SHIP_VY
	ship.direction = INITIAL_DIRECTION_DEGREES
}

func (ship *Ship) Update() {}
