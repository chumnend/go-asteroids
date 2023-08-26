package asteroids

const (
	SHIP_START_X      = GAME_WIDTH / 2
	SHIP_START_Y      = GAME_HEIGHT / 2
	INITIAL_DIRECTION = 0 // radians
	INITIAL_SHIP_VX   = 0
	INITIAL_SHIP_VY   = 0
	TURN_RATE         = 5
)

type Ship struct {
	Entity
	Direction float64
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

	// set initial player position
	ship.position.X = SHIP_START_X
	ship.position.Y = SHIP_START_Y
	ship.Direction = INITIAL_DIRECTION
	ship.velocity.X = INITIAL_SHIP_VX
	ship.velocity.Y = INITIAL_SHIP_VY

	return ship, nil
}
