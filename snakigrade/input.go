package snakigrade

type mouseState int

// Dir represents a direction.
type Dir int

const (
	dirUp Dir = iota
	dirRight
	dirDown
	dirLeft
)

type touchState int

const (
	touchStateNone touchState = iota
	touchStatePressing
	touchStateSettled
	touchStateInvalid
)

// Input represents the current key states.
type Input struct {
	mouseState    mouseState
	mouseInitPosX int
	mouseInitPosY int
	mouseDir      Dir

	touchState    touchState
	touchID       int
	touchInitPosX int
	touchInitPosY int
	touchLastPosX int
	touchLastPosY int
	touchDir      Dir
}
