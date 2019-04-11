package snakigrade

import (
	"github.com/hajimehoshi/ebiten"
)

// Board represents the game board.
type Board struct {
	size  int
	tiles map[*Tile]struct{}
	tasks []task
}
