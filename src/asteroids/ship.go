package asteroids

import "math"

const (
	INITIAL_SHIP_X            = GAME_WIDTH / 2
	INITIAL_SHIP_Y            = GAME_HEIGHT / 2
	INITIAL_SHIP_VX           = 0
	INITIAL_SHIP_VY           = 0
	INITIAL_DIRECTION_DEGREES = 0
	MAX_SPEED                 = 1.8
	ACCELERATION              = 0.1
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

func (ship *Ship) Accelerate(opposite bool) {
	// calculate velocity vector based on direction
	sign := 1.
	if opposite {
		sign *= -1
	}

	v1 := math.Sqrt(math.Pow(ship.Velocity.X, 2) + math.Pow(ship.Velocity.Y, 2))
	v2 := v1 + sign*ACCELERATION

	dirRad := degreeToRad(ship.Direction - 90) // assuming direction 0 deg is UP, need to rotate -ve
	vx := v2 * math.Cos(dirRad)
	vy := v2 * math.Sin(dirRad)

	if vx > MAX_SPEED {
		vx = MAX_SPEED
	}

	if vx < -MAX_SPEED {
		vx = -MAX_SPEED
	}

	if vy > MAX_SPEED {
		vy = MAX_SPEED
	}

	if vy < -MAX_SPEED {
		vy = -MAX_SPEED
	}

	ship.Velocity.X = vx
	ship.Velocity.Y = vy
}

func (ship *Ship) Rotate(clockwise bool) {
	sign := 1.
	if !clockwise {
		sign *= -1
	}

	newDir := ship.Direction + sign*float64(TURN_RATE)
	newDir = math.Mod(newDir, 360)
	if newDir < 0 {
		newDir += 360
	}

	ship.Direction = newDir
}

func (ship *Ship) Initialize() {
	ship.Position.X = INITIAL_SHIP_X
	ship.Position.Y = INITIAL_SHIP_Y
	ship.Velocity.X = INITIAL_SHIP_VX
	ship.Velocity.Y = INITIAL_SHIP_VY
	ship.Direction = INITIAL_DIRECTION_DEGREES
}

func (ship *Ship) Update() {
	ship.Position.X += ship.Velocity.X
	if ship.Position.X < 0 {
		ship.Position.X = 0
	}
	if ship.Position.X > GAME_WIDTH {
		ship.Position.X = GAME_WIDTH
	}

	ship.Position.Y += ship.Velocity.Y
	if ship.Position.Y < 0 {
		ship.Position.Y = 0
	}
	if ship.Position.Y > GAME_HEIGHT {
		ship.Position.Y = GAME_HEIGHT
	}
}
