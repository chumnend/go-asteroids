package asteroids

import (
	"github.com/hajimehoshi/ebiten/v2"
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

	switch g.gameState {
	case GameStatePlaying:
		for _, a := range g.asteroids {
			a.updatePosition()
		}
	}

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

// pauseGame will halt all action on screen and show the pause menu
// TODO: logic for pausing game
func (g *Game) pauseGame() {
	g.showMenu(MenuPause)
}
