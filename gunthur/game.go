package gunthur

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 320
	screenHeight = 256
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
	count int
	state GameState

	player1 *Player
	player2 *Player
	player3 *Player

	background *ebiten.Image
}

func (g *Game) Init() error {
	var err error
	g.background, _, err = ebitenutil.NewImageFromFile("assets/backgrounds/bg_320x256.png")
	if err != nil {
		return err
	}

	g.player1 = NewPlayer(50, 100)
	g.player2 = NewPlayer(100, 100)
	g.player3 = NewPlayer(150, 100)

	return nil
}

func (g *Game) Update() error {
	g.count++

	switch g.state {
	case GameStateMenu:
		if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
			g.state = GameStatePlaying
		}
	case GameStatePlaying:
		if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
			g.state = GameStateMenu
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.background, nil)

	g.player1.idle(screen, g.count)
	g.player2.runRight(screen, g.count)
	g.player3.jump(screen, g.count)

	// debug message
	ebitenutil.DebugPrint(screen, "Current state: "+strconv.Itoa(int(g.state))+" Frame: "+strconv.Itoa(int(g.count)))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
