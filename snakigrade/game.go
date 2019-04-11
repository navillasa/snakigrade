package snakigrade

import (
	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 420
	screenHeight = 600
	boardSize    = 4
)

// Game represents a game state.
type Game struct {
	input      *Input
	board      *Board
	boardImage *ebiten.Image
}
