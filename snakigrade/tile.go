package snakigrade

// TileData represents a tile information like a value and a position.
type TileData struct {
	value int
	x     int
	y     int
}

// Tile represents a tile information including TileData and animation states.
type Tile struct {
	current TileData

	// next represents a next tile information after moving.
	// next is empty when the tile is not about to move.
	next TileData

	movingCount       int
	startPoppingCount int
	poppingCount      int
}
