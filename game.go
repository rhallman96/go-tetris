package main

const (
	BoardHeight = 20
	BoardWidth  = 10
)

// On the game board, any value other than zero indicates a block is present.
// Depending on how the GUI is implemented, this corresponds to color/pattern/etc.
type BlockValue uint8

type Shape []Point

type Point struct {
	X, Y int
}

type Piece struct {
	Value BlockValue

	coords        Point
	rotationIndex int
	rotations     []Shape
}

type Board struct {
	Grid  [][]BlockValue
	Piece Piece
}

var pieceRotations = [][]Shape{

	// Square
	Shape{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	}.identityRotations(),

	// Line
	Shape{
		{1, 0},
		{1, 1},
		{1, 2},
		{1, 3},
	}.flipRotations(),

	// T-Shape
	Shape{
		{0, 1},
		{1, 1},
		{2, 1},
		{1, 2},
	}.squareRotations(),

	// L-Shape
	Shape{
		{1, 0},
		{1, 1},
		{1, 2},
		{2, 2},
	}.squareRotations(),

	// Flipped L-Shape
	Shape{
		{1, 0},
		{1, 1},
		{1, 2},
		{0, 2},
	}.squareRotations(),
}

// Reset clears the board and selects the initial piece.
func (board *Board) Reset() {
	board.Grid = make([][]BlockValue, BoardHeight)
	for i := 0; i < BoardHeight; i++ {
		board.Grid[i] = make([]BlockValue, BoardWidth)
	}
}

// ClearFilledRows removes filled rows from the Board.
func (board *Board) ClearFilledRows() {
	rows := board.FilledRowIndices()
	if len(rows) == 0 {
		return
	}

	for _, v := range rows {
		board.clearRow(v)
	}
}

// FilledRowIndices returns a slice containing the indices of filled rows
// on the Board. It does not clear these rows or modify the Board's state.
func (board *Board) FilledRowIndices() []int {
	var rows []int
	for i, _ := range board.Grid {
		if board.isRowFilled(i) {
			rows = append(rows, i)
		}
	}
	return rows
}

func (board *Board) isRowFilled(index int) bool {
	row := board.Grid[index]
	for _, v := range row {
		if v == 0 {
			return false
		}
	}
	return true
}

func (board *Board) clearRow(index int) bool {
	for i := index - 1; i >= 0; i-- {
		board.Grid[i+1] = board.Grid[i]
	}
	board.Grid[0] = make([]BlockValue, BoardWidth)
	return true
}

func (piece *Piece) GetShape() Shape {
	if len(piece.rotations) == 0 {
		return Shape{}
	}

	shape := piece.rotations[piece.rotationIndex%len(piece.rotations)]
	for i, _ := range shape {
		shape[i].X += piece.coords.X
		shape[i].Y += piece.coords.Y
	}
	return shape
}

func (piece *Piece) RotateLeft() bool {
	return true
}

func (piece *Piece) RotateRight() bool {
	return true
}

func (shape Shape) squareRotations() []Shape {
	transforms := []Shape{shape}
	for i := 0; i < 3; i++ {
		for j, point := range shape {
			x := point.X
			point.X = point.Y
			point.Y = -x
			shape[j] = point
		}
		transforms = append(transforms, shape)
	}

	return transforms
}

func (shape Shape) flipRotations() []Shape {
	transforms := []Shape{shape}
	flipped := shape
	for i, _ := range flipped {
		flipped[i].X = shape[i].Y
		flipped[i].Y = shape[i].X
	}
	transforms = append(transforms, flipped)
	return transforms
}

func (shape Shape) identityRotations() []Shape {
	return []Shape{shape}
}
