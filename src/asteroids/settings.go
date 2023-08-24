package asteroids

import "image/color"

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
	MenuMain MenuState = iota
	MenuPause
	MenuGameOver
	MenuWin
)

var (
	DPI           = 72
	FONT_SIZE     = 8
	MENU_BG_COLOR = color.RGBA{72, 170, 182, 200}
	TEXT_COLOR    = color.RGBA{34, 32, 32, 255}
)

type textItem struct {
	text  string
	posY  int
	color color.Color
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
