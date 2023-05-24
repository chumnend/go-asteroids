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
	count      int
	state      GameState
	player     *Player
	background *ebiten.Image
}

func (g *Game) Init() error {
	var err error
	g.background, _, err = ebitenutil.NewImageFromFile("assets/backgrounds/bg_320x256.png")
	if err != nil {
		return err
	}

	g.player = NewPlayer(screenWidth/2, screenHeight/2)

	return nil
}

func (g *Game) Update() error {
	g.count++

	switch g.state {
	case GameStateMenu:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.state = GameStatePlaying
		}
	case GameStatePlaying:
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
			g.player.state = PlayerStateJumping
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
			g.player.state = PlayerStateCrouch
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
			g.player.state = PlayerStateRunRight
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
			g.player.state = PlayerStateRunLeft
		} else if inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) || inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) || inpututil.IsKeyJustReleased(ebiten.KeyArrowUp) || inpututil.IsKeyJustReleased(ebiten.KeyArrowDown) {
			g.player.state = PlayerStateIdle
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.background, nil)

	g.player.draw(screen, g.count)

	// debug message
	ebitenutil.DebugPrint(screen, "Current state: "+strconv.Itoa(int(g.state))+" Frame: "+strconv.Itoa(int(g.count)))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
