package asteroids

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
)

const (
	// internal game size
	gameWidth  = 300
	gameHeight = 250

	scale        = 2 // scale 200% in window
	windowWidth  = gameWidth * scale
	windowHeight = gameHeight * scale

	dt = 1 / 60.0 // assume that delta is fixed and we're always running at 60 FPS
)

type GameState int

const (
	GameStateMenu GameState = iota
	GameStatePlaying
)

// Game implements ebiten.Game interface
type Game struct {
	ship      *Ship
	asteroids []*Asteroid

	gameState GameState
	menuState MenuState

	font font.Face

	pressedKeys []ebiten.Key
}

// NewGame returns a Game struct, takes the size of the game screen
func NewGame() (*Game, int, int) {
	return &Game{}, windowWidth, windowHeight
}

// Init loads all resources for the game
func (g *Game) Init() error {
	if err := g.loadMenuResources(); err != nil {
		return err
	}

	if err := g.loadObjects(); err != nil {
		return err
	}

	g.showMenu(MenuMain)

	return nil
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.processInput()
	g.updateAsteroids()

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	switch g.gameState {
	case GameStateMenu:
		g.drawMenu(screen)
	case GameStatePlaying:
		g.ship.Draw(screen)
		for _, a := range g.asteroids {
			a.Draw(screen)
		}
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return gameWidth, gameHeight
}

// === HELPER FUNCTIONS ================================================================================

// loadObjects should only be called before the game is run, and is responsible for loading all game objects
func (g *Game) loadObjects() error {
	// load ship
	ship, err := makeShip()
	if err != nil {
		return err
	}
	g.ship = ship

	// load asteroids
	asteroids, err := makeAsteroids()
	if err != nil {
		return err
	}
	g.asteroids = asteroids

	return nil
}

// showMenu will change the game state to show menus and show the menu state that is given
func (g *Game) showMenu(state MenuState) {
	g.gameState = GameStateMenu
	g.menuState = state
}

// startGame changes the game state to playing
// TODO: logic to start game from current state
func (g *Game) startGame() {
	g.gameState = GameStatePlaying
}

// restartGame resets the location of all objects
func (g *Game) restartGame() {
	// reset the ships position
	g.ship.resetPosition()

	// reset asteroids
	for _, a := range g.asteroids {
		a.resetPosition()
	}

	g.gameState = GameStatePlaying
}

// pauseGame will halt all action on screen and show the pause menu
// TODO: logic for pausing game
func (g *Game) pauseGame() {
	g.showMenu(MenuPause)
}

// checkCollision checks if two entites overlap
func (g *Game) checkCollision(o1 *Entity, o2 *Entity) bool {
	rect1 := o1.getAABB().ToImageRect()
	rect2 := o2.getAABB().ToImageRect()

	return rect1.Overlaps(rect2)
}

// processInput reads user input and updates game accordingly
func (g *Game) processInput() {
	switch g.gameState {
	case GameStateMenu:
		switch g.menuState {
		case MenuMain:
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
				g.startGame()
			}
		case MenuPause:
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
				g.startGame()
			}
		case MenuGameOver:
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
				g.restartGame()
			}
		case MenuWin:
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
				g.restartGame()
			}
		default:
			panic("unexpected state")
		}

	case GameStatePlaying:
		if inpututil.IsKeyJustPressed(ebiten.KeyP) {
			g.pauseGame()
		}

		g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])

		for _, key := range g.pressedKeys {
			switch key {
			case ebiten.KeyW:
				g.ship.moveUp()
			case ebiten.KeyS:
				g.ship.moveDown()
			case ebiten.KeyA:
				g.ship.moveLeft()
			case ebiten.KeyD:
				g.ship.moveRight()
			case ebiten.KeyQ:
				g.ship.rotateLeft()
			case ebiten.KeyE:
				g.ship.rotateRight()
			}
		}
	}
}

// updateAsteroids updates the positions for objects on screen
func (g *Game) updateAsteroids() {
	switch g.gameState {
	case GameStatePlaying:
		collided := false

		for idx, a := range g.asteroids {
			a.updatePosition()

			// check if any asteroids collided with the ship
			didCollide := g.checkCollision(&a.Entity, &g.ship.Entity)
			if didCollide {
				collided = true
			}

			// check if asteroid collided with other asteroids
			otherAsteroids := make([]*Asteroid, len(g.asteroids))
			copy(otherAsteroids, g.asteroids)
			otherAsteroids = append(otherAsteroids[:idx], otherAsteroids[idx+1:]...)
			for _, oA := range otherAsteroids {
				didBounce := g.checkCollision(&a.Entity, &oA.Entity)
				if didBounce {
					a.bounce()
				}
			}
		}

		// on ship collision go to game over state
		if collided {
			g.showMenu(MenuGameOver)
		}
	}
}
