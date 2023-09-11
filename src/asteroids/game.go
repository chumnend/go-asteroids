package asteroids

import (
	"errors"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
//
//    _____    _________________________________________ ________  .___________    _________
//  /  _  \  /   _____/\__    ___/\_   _____/\______   \\_____  \ |   \______ \  /   _____/
// /  /_\  \ \_____  \   |    |    |    __)_  |       _/ /   |   \|   ||    |  \ \_____  \
///    |    \/        \  |    |    |        \ |    |   \/    |    \   ||    `   \/        \
//\____|__  /_______  /  |____|   /_______  / |____|_  /\_______  /___/_______  /_______  /
//        \/        \/                    \/         \/         \/            \/        \/
//
// By Nicholas Chumney (go clone of the classic game Asteroids)
//
///////////////////////////////////////////////////////////////////////////////////////////////////

// This file conatins 6 sections:
//   1) SETTINGS - constants and structs used in the game
//   2) MAIN - the main ebiten game struct
//   3) INPUT - for reading and handling keyboard inputs
//   4) PROCESS - for updating game state
//   5) PAINT - for drawing to the screen

// 1) SETTINGS =======================================================================================

const (
	GAME_WIDTH        = 300 // internal game width
	GAME_HEIGHT       = 300 // internal game height
	SCALE             = 2   // factor to scale for desktop
	WINDOW_WIDTH      = GAME_WIDTH * SCALE
	WINDOW_HEIGHT     = GAME_HEIGHT * SCALE
	DT                = 1 / 60.0 // assume the delta is fixed and we are always at 60 FPS
	DPI               = 72
	FONT_SIZE         = 8
	AUDIO_SAMPLE_RATE = 48000
)

var (
	MENU_BG_COLOR = color.RGBA{72, 170, 182, 200}
	TEXT_COLOR    = color.RGBA{34, 32, 32, 255}
)

type GameState int

const (
	GameStateMenu GameState = iota
	GameStatePlaying
)

type MenuState int

const (
	MenuStateMain MenuState = iota
	MenuStatePause
	MenuStateGameOver
	MenuStateWin
)

type textItem struct {
	text   string
	height int
	color  color.Color
}

var mainMenuTexts = []textItem{
	{"Ebiten Asteroids", GAME_HEIGHT / 2, TEXT_COLOR},
	{"Press Space to start", GAME_HEIGHT/2 + FONT_SIZE*2, TEXT_COLOR},
}

var pauseTexts = []textItem{
	{"PAUSED", GAME_HEIGHT / 2, TEXT_COLOR},
	{"Press Space to continue", GAME_HEIGHT/2 + FONT_SIZE*2, TEXT_COLOR},
}

var gameOverTexts = []textItem{
	{"Game Over", GAME_HEIGHT / 2, TEXT_COLOR},
	{"Press Space to replay", GAME_HEIGHT/2 + FONT_SIZE*2, TEXT_COLOR},
}

var wonTexts = []textItem{
	{"You Won!", GAME_HEIGHT / 2, TEXT_COLOR},
	{"Press Space to replay", GAME_HEIGHT/2 + FONT_SIZE*2, TEXT_COLOR},
}

// 2) MAIN ===========================================================================================

// Game implements the ebiten.Game interface
type Game struct {
	gameState   GameState
	menuState   MenuState
	pressedKeys []ebiten.Key
	ship        *Ship
	asteroids   Asteroids
	bullets     Bullets
	showDebug   bool

	font font.Face
}

// NewGame returns a Game struct, the width of the window and the height of the window
func NewGame() (*Game, int, int) {
	return &Game{
		gameState:   GameStateMenu,
		menuState:   MenuStateMain,
		pressedKeys: nil,
		ship:        nil,
		asteroids:   nil,
		bullets:     nil,
		font:        nil,
		showDebug:   false,
	}, WINDOW_WIDTH, WINDOW_HEIGHT
}

// Init loads all resources for the game
func (game *Game) Init() error {
	// load the font type
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		return err
	}
	tf, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(FONT_SIZE),
		DPI:     float64(DPI),
		Hinting: font.HintingFull,
	})
	if err != nil {
		return err
	}
	game.font = tf

	// load ship
	ship, err := NewShip()
	if err != nil {
		return err
	}
	game.ship = ship

	// load asteroids
	asteroids, err := NewAsteroids(ship)
	if err != nil {
		return err
	}
	game.asteroids = asteroids

	// load bullet
	bullets, err := NewBullets()
	if err != nil {
		return err
	}
	game.bullets = bullets

	return nil
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (game *Game) Update() error {
	if err := game.handleInput(); err != nil {
		return err
	}

	if err := game.processLogic(); err != nil {
		return err
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (game *Game) Draw(screen *ebiten.Image) {
	switch game.gameState {
	case GameStateMenu:
		game.drawMenuScreen(screen)
	case GameStatePlaying:
		game.drawObjects(screen)
	}

	// debug info
	if game.showDebug {
		game.printDebugInfo(screen)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return GAME_WIDTH, GAME_HEIGHT
}

// 3) INPUT ===================================================================================

// handleInput reads key inputs and performs actions
// WS - to speed up/down
// AD - to rotate left/right
// SPACE (in game) - to shoot
// SPACE (in menus) - to starts game
// P - to pause
// 0 - show debug info
// ESC - quit
func (game *Game) handleInput() error {
	// get pressed keys
	game.pressedKeys = inpututil.AppendPressedKeys(game.pressedKeys[:0])

	// force game end
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	// show debug info
	if inpututil.IsKeyJustPressed(ebiten.Key0) {
		game.showDebug = !game.showDebug
	}

	switch game.gameState {
	case GameStateMenu:
		switch game.menuState {
		case MenuStateMain:
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
				game.startGame()
			}
		case MenuStatePause:
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
				game.resumeGame()
			}
		case MenuStateGameOver:
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
				game.restartGame()
			}
		case MenuStateWin:
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
				game.restartGame()
			}
		default:
			return errors.New(fmt.Sprintf("no input for menu %v was found", game.menuState))
		}
	case GameStatePlaying:
		if inpututil.IsKeyJustPressed(ebiten.KeyP) {
			game.pauseGame()
		}

		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			game.bullets.Shoot(game.ship)
		}

		for _, key := range game.pressedKeys {
			switch key {
			case ebiten.KeyW:
				game.ship.Accelerate(false)
			case ebiten.KeyS:
				game.ship.Accelerate(true)
			case ebiten.KeyA:
				game.ship.Rotate(false) // rotate counter clockwise
			case ebiten.KeyD:
				game.ship.Rotate(true) // rotate clockwise
			}
		}
	default:
		return errors.New(fmt.Sprintf("no input for %v was found", game.gameState))
	}

	return nil
}

// startGame sets the game state to "playing"
func (game *Game) startGame() {
	game.gameState = GameStatePlaying
}

// pauseGame sets the game state to "menu" and show the "pause" menu
func (game *Game) pauseGame() {
	game.gameState = GameStateMenu
	game.menuState = MenuStatePause
}

// resumeGame sets the game state to "playing" and hides the "pause" menu
func (game *Game) resumeGame() {
	game.gameState = GameStatePlaying
}

// restartGame sets the game state to "menu" and show the "main" menu
func (game *Game) restartGame() {
	// reset game object parameters
	game.ship.Initialize()
	game.asteroids.Initialize(game.ship)
	game.bullets.Initialize()
	game.gameState = GameStateMenu
	game.menuState = MenuStateMain
}

// winGame sets the game state to "menu" and show the "win" menu
func (game *Game) winGame() {
	game.gameState = GameStateMenu
	game.menuState = MenuStateWin
}

// loseGame sets the game state to "menu" and shows the "game over" menu
func (game *Game) loseGame() {
	game.gameState = GameStateMenu
	game.menuState = MenuStateGameOver
}

// 4) PROCESS =============================================================================

// checkCollisions checks for any objects that are colliding
func (game *Game) checkCollisions() {
	// if bullet collides with asteroids, destroy the asteroid
	for _, bullet := range game.bullets {
		for _, asteroid := range game.asteroids {
			if asteroid.CollidesWith(&bullet.Entity) {
				bullet.Initialize() // reset the bullet
				asteroid.Destroy()  // remove the asteroid
			}
		}
	}

	// on ship collision with any asteroids, end the game, show LOSE screen
	for idx, asteroid := range game.asteroids {
		if game.ship.CollidesWith(&asteroid.Entity) {
			game.loseGame()
		}

		// check if asteroid collided with other asteroids
		otherAsteroids := make(Asteroids, len(game.asteroids))
		copy(otherAsteroids, game.asteroids)
		otherAsteroids = append(otherAsteroids[:idx], otherAsteroids[idx+1:]...)
		for _, oA := range otherAsteroids {
			didBounce := asteroid.CollidesWith(&oA.Entity)
			if didBounce {
				asteroid.Bounce()
			}
		}

	}
}

// checkWin checks if player has won the game (all asteroids destroyed)
func (game *Game) checkWin() bool {
	didWin := true
	for _, asteroid := range game.asteroids {
		if asteroid.IsHidden == false {
			didWin = false
		}
	}
	return didWin
}

// processLogic updates all game objects each frame
func (game *Game) processLogic() error {
	if game.gameState == GameStatePlaying {
		// update game objects
		game.ship.Update()
		game.asteroids.Update()
		game.bullets.Update()

		// evaluate state of the game
		game.checkCollisions()
		if game.checkWin() == true {
			game.winGame()
		}
	}

	return nil
}

// 5) PAINT  ===================================================================================

// drawMenuScreen draws the text to be shown on menu based on current menu state
func (g *Game) drawMenuScreen(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, 0, 0, GAME_WIDTH, GAME_HEIGHT, MENU_BG_COLOR)

	// get the text items based on menu state
	var texts []textItem
	switch g.menuState {
	case MenuStateMain:
		texts = mainMenuTexts
	case MenuStatePause:
		texts = pauseTexts
	case MenuStateGameOver:
		texts = gameOverTexts
	case MenuStateWin:
		texts = wonTexts
	default:
		panic("unexpected state")
	}

	// draw each text item to the screen
	for _, ti := range texts {
		text.Draw(screen, ti.text, g.font, GAME_WIDTH/2.-len(ti.text)/2.*FONT_SIZE, ti.height, ti.color)
	}
}

// drawObjects paints the screen with game objects
func (game *Game) drawObjects(screen *ebiten.Image) {
	game.ship.Draw(screen)
	game.asteroids.Draw(screen)
	game.bullets.Draw(screen)
}

// printDebugInfo shows game information if debug mode is on
func (game *Game) printDebugInfo(screen *ebiten.Image) {
	var currentGameState string
	switch game.gameState {
	case GameStateMenu:
		currentGameState = "Menu"
	case GameStatePlaying:
		currentGameState = "Playing"
	}

	var currentMenuState string
	switch game.menuState {
	case MenuStateMain:
		currentMenuState = "Main Menu"
	case MenuStatePause:
		currentMenuState = "Paused"
	case MenuStateGameOver:
		currentMenuState = "Game Over"
	case MenuStateWin:
		currentMenuState = "Win"
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("State: %v   Menu: %v\nKeys: %v\nShip Pos X:%.0f Y:%.0f Dir: %.1f\nShip Spd Vx: %.1f Vy: %.1f ", currentGameState, currentMenuState, game.pressedKeys, game.ship.Position.X, game.ship.Position.Y, game.ship.Velocity.X, game.ship.Velocity.Y, game.ship.Direction))
}
