package gunthur

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 320
	screenHeight = 240
	scale        = 2

	WindowWidth  = screenWidth * scale
	WindowHeight = screenHeight * scale
)

type GameState int

const (
	GameStateMenu = iota
	GameStatePlaying
)

type Game struct {
	state  GameState
	player *Sprite
}

func NewGame() *Game {
	g := &Game{}
	g.state = GameStatePlaying
	playerSpritesheet := NewSpritesheet("./assets/sprites/adventurer/adventurer-sheet.png", 50, 37)
	g.player = NewSprite(playerSpritesheet)
	g.player.AddAnimation("idle", 1200, 0, 0, 4, false, false)
	g.player.AddAnimation("runRight", 1000, 50, 37, 6, false, false)
	g.player.AddAnimation("runLeft", 1000, 50, 37, 6, false, true)
	g.player.AddAnimation("jump", 600, 0, 74, 10, false, false)
	g.player.AddAnimation("crouch", 1500, 200, 0, 4, false, false)
	g.player.SetAnimation("idle")
	return g
}

func (g *Game) Update() error {
	switch g.state {
	case GameStateMenu:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.state = GameStatePlaying
		}
	case GameStatePlaying:
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
			g.player.SetAnimation("jump")
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
			g.player.SetAnimation("crouch")
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
			g.player.SetAnimation("runRight")
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
			g.player.SetAnimation("runLeft")
		} else if inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) || inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) || inpututil.IsKeyJustReleased(ebiten.KeyArrowUp) || inpututil.IsKeyJustReleased(ebiten.KeyArrowDown) {
			g.player.SetAnimation("idle")
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "This is a test.")
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
