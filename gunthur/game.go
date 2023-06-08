package gunthur

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameState int

const (
	GameStateMenu = iota
	GameStatePlaying
	GameStatePaused
)

type Game struct {
	state        GameState
	player       *Sprite
	screenWidth  int
	screenHeight int
	pressedKeys  []ebiten.Key
}

func NewGame(width int, height int) *Game {
	return &Game{
		screenWidth:  width,
		screenHeight: height,
	}
}

func (g *Game) Init() error {
	g.state = GameStatePlaying
	playerSpritesheet := NewSpritesheet("./assets/sprites/adventurer/adventurer-sheet.png", 50, 37)
	g.player = NewSprite(playerSpritesheet)
	g.player.AddAnimation("idle", 1000, 0, 0, 4, false, false)
	g.player.AddAnimation("runRight", 1000, 50, 37, 6, false, false)
	g.player.AddAnimation("runLeft", 1000, 50, 37, 6, false, true)
	g.player.AddAnimation("jump", 1500, 0, 74, 10, false, false)
	g.player.AddAnimation("crouch", 1000, 200, 0, 4, false, false)
	g.player.SetAnimation("idle")
	return nil
}

func (g *Game) Update() error {
	g.player.UpdatePosition()

	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])
	for _, key := range g.pressedKeys {
		switch key.String() {
		case "ArrowDown":
			g.player.SetAnimation("crouch")
		case "ArrowUp":
			g.player.SetAnimation("jump")
		case "ArrowRight":
			g.player.SetAnimation("runRight")
			g.player.SetSpeed(2, 0)
		case "ArrowLeft":
			g.player.SetAnimation("runLeft")
			g.player.SetSpeed(-2, 0)
		}
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) || inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) || inpututil.IsKeyJustReleased(ebiten.KeyArrowUp) || inpututil.IsKeyJustReleased(ebiten.KeyArrowDown) {
		g.player.SetAnimation("idle")
		g.player.SetSpeed(0, 0)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "This is a test.")
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}
