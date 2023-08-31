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

// SETTINGS =======================================================================================

const (
	GAME_WIDTH    = 300 // internal game width
	GAME_HEIGHT   = 300 // internal game height
	SCALE         = 2   // factor to scale for desktop
	WINDOW_WIDTH  = GAME_WIDTH * SCALE
	WINDOW_HEIGHT = GAME_HEIGHT * SCALE
	DT            = 1 / 60.0 // assume the delta is fixed and we are always at 60 FPS
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

var (
	DPI           = 72
	FONT_SIZE     = 8
	MENU_BG_COLOR = color.RGBA{72, 170, 182, 200}
	TEXT_COLOR    = color.RGBA{34, 32, 32, 255}
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

// MAIN ===========================================================================================

// Game implements the ebiten.Game interface
type Game struct {
	gameState   GameState
	menuState   MenuState
	pressedKeys []ebiten.Key

	font font.Face

	ship     *Ship
	asteroid *Asteroid
}

// NewGame returns a Game struct, the width of the window and the height of the window
func NewGame() (*Game, int, int) {
	return &Game{}, WINDOW_WIDTH, WINDOW_HEIGHT
}

// Init loads all resources for the game
func (game *Game) Init() error {
	if err := game.loadMenuResources(); err != nil {
		return err
	}

	if err := game.loadObjects(); err != nil {
		return err
	}

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
	game.printDebugInfo(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return GAME_WIDTH, GAME_HEIGHT
}

// LOADING ========================================================================================

// loadObjects loads all required assets for the game
func (game *Game) loadObjects() error {
	// load ship
	ship, err := NewShip()
	if err != nil {
		return err
	}
	game.ship = ship

	// load asteroid
	asteroid, err := NewAsteroid()
	if err != nil {
		return err
	}
	// make sure asteroid does not spawn on ship initially
	for asteroid.CollidesWith(&ship.Entity) {
		asteroid.GetRandomPosition()
	}
	game.asteroid = asteroid

	return nil
}

func (game *Game) loadMenuResources() error {
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

	return nil
}

// HANDLE INPUT ===================================================================================

// handleInput reads key inputs and performs actions
func (game *Game) handleInput() error {
	// get pressed keys
	game.pressedKeys = inpututil.AppendPressedKeys(game.pressedKeys[:0])

	// force game end
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
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
		default:
			return errors.New("unexpected menu state")
		}
	case GameStatePlaying:
		if inpututil.IsKeyJustPressed(ebiten.KeyP) {
			game.pauseGame()
		}

		for _, key := range game.pressedKeys {
			switch key {
			case ebiten.KeyW:
				game.ship.Accelerate()
			case ebiten.KeyS:
				game.ship.Decelerate()
			case ebiten.KeyA:
				game.ship.Rotate(false) // rotate counter clockwise
			case ebiten.KeyD:
				game.ship.Rotate(true) // rotate clockwise
			}
		}
	default:
		return errors.New("unexpected game state")
	}

	return nil
}

func (game *Game) startGame() {
	game.gameState = GameStatePlaying
}

func (game *Game) pauseGame() {
	game.gameState = GameStateMenu
	game.menuState = MenuStatePause
}

func (game *Game) resumeGame() {
	game.gameState = GameStatePlaying
}

func (game *Game) restartGame() {
	// reset game object parameters
	game.ship.Initialize()
	game.asteroid.Initialize()

	game.gameState = GameStateMenu
	game.menuState = MenuStateMain
}

func (game *Game) gameOver() {
	game.gameState = GameStateMenu
	game.menuState = MenuStateGameOver
}

// PROCESS GAME LOGIC =============================================================================

// checkCollisions checks for any objects that are colliding
func (game *Game) checkCollisions() {
	// on ship collision with any asteroids, end the game
	if game.ship.CollidesWith(&game.asteroid.Entity) {
		game.gameOver()
	}

}

// processLogic updates all game objects each frame
func (game *Game) processLogic() error {
	if game.gameState == GameStatePlaying {
		game.ship.Update()
		game.asteroid.Update()
		game.checkCollisions()
	}

	return nil
}

// PAINT SCREEN ===================================================================================

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

func (game *Game) drawObjects(screen *ebiten.Image) {
	game.ship.Draw(screen)
	game.asteroid.Draw(screen)
}

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

	ebitenutil.DebugPrint(screen, fmt.Sprintf("State: %v   Menu: %v\nKeys: %v", currentGameState, currentMenuState, game.pressedKeys))
}
