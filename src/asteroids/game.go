package asteroids

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GameState int

const (
	GameStateMenu = iota
	GameStatePlaying
	GameStateGameOver
)

var (
	screenWidth  int
	screenHeight int
)

// Game implements ebiten.Game interface
type Game struct {
	state        GameState
	pressedKeys  []ebiten.Key
	currentLevel *Scene
	levels       map[int]*Scene
}

// NewGame returns a Game struct, takes the size of the game screen
func NewGame(width, height int) *Game {
	levels := make(map[int]*Scene)
	level1 := NewScene()
	level1.AddComponent(NewPlayer())
	level1.AddComponent(NewObstacle())
	levels[1] = level1

	screenWidth = width
	screenHeight = height

	return &Game{
		currentLevel: levels[1],
		levels:       levels,
	}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.processInput()
	g.currentLevel.Update(g.pressedKeys)
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case GameStateMenu:
		g.drawStartMenu(screen)
	case GameStatePlaying:
		g.currentLevel.Draw(screen, ebiten.DrawImageOptions{})
	case GameStateGameOver:
		g.drawGameOver(screen)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
