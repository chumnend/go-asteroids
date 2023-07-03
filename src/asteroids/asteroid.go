package asteroids

type Asteroid struct {
	Entity
}

func makeAsteroid() (*Asteroid, error) {
	return &Asteroid{}, nil
}
