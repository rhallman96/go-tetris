package game

import (
	"math/rand"
	"time"
)

const (
	BoardHeight = 20
	BoardWidth  = 11
)

// Init function sets seed for random piece generation.
func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type Shape []Point

type Point struct {
	X, Y int
}

type Piece struct {
	Value int

	coords        Point
	rotationIndex int
	rotations     []Shape
}

// Rotations are pre-computed for each unique kind of piece.
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

type Board struct {
	Grid  [][]int
	Piece Piece
}

// Reset clears the board and selects the initial piece.
func (board *Board) Reset() {
	board.Grid = make([][]int, BoardHeight)
	for i := 0; i < BoardHeight; i++ {
		board.Grid[i] = make([]int, BoardWidth)
	}
	board.NextPiece()
}

func (board *Board) NextPiece() bool {
	blockValue := int(rand.Intn(256))
	rotations := pieceRotations[rand.Intn(len(pieceRotations))]
	coords := Point{(BoardWidth / 2) - 1, 0}

	board.Piece = Piece{blockValue, coords, 0, rotations}
	if board.pieceOverlapsGrid() {
		return false
	}
	return true
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

func (board *Board) DropPiece() bool {
	board.Piece.coords.Y++
	if !board.pieceOverlapsGrid() {
		return true
	}

	board.Piece.coords.Y--
	shape := board.Piece.GetShape()

	for _, point := range shape {
		board.Grid[point.Y][point.X] = board.Piece.Value
	}

	return false
}

func (board *Board) MovePieceLeft() bool {
	board.Piece.coords.X--
	if board.pieceOverlapsGrid() {
		board.Piece.coords.X++
		return false
	}
	return true
}

func (board *Board) MovePieceRight() bool {
	board.Piece.coords.X++
	if board.pieceOverlapsGrid() {
		board.Piece.coords.X--
		return false
	}
	return true
}

func (board *Board) RotatePieceLeft() bool {
	board.Piece.rotationIndex--
	if board.Piece.rotationIndex < 0 {
		board.Piece.rotationIndex += len(board.Piece.rotations)
	}

	if board.pieceOverlapsGrid() {
		board.Piece.rotationIndex++
		return false
	}
	return true
}

func (board *Board) RotatePieceRight() bool {
	board.Piece.rotationIndex++
	if board.pieceOverlapsGrid() {
		board.Piece.rotationIndex--
		return false
	}
	return true
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
	board.Grid[0] = make([]int, BoardWidth)
	return true
}

func (board *Board) pieceOverlapsGrid() bool {
	shape := board.Piece.GetShape()
	for _, point := range shape {
		if point.X < 0 || point.X >= BoardWidth ||
			point.Y < 0 || point.Y >= BoardHeight ||
			board.Grid[point.Y][point.X] != 0 {
			return true
		}
	}
	return false
}

func (piece *Piece) GetShape() Shape {
	if len(piece.rotations) == 0 {
		return Shape{}
	}

	rotation := piece.rotations[piece.rotationIndex%len(piece.rotations)]
	shape := make(Shape, len(rotation))
	copy(shape, rotation)

	for i, _ := range shape {
		shape[i].X += piece.coords.X
		shape[i].Y += piece.coords.Y
	}
	return shape
}

func (shape Shape) squareRotations() []Shape {
	transforms := []Shape{shape}
	for i := 0; i < 3; i++ {
		next := make(Shape, len(shape))
		for j, point := range transforms[i] {
			next[j] = Point{2 - point.Y, point.X}
		}
		transforms = append(transforms, next)
	}

	return transforms
}

func (shape Shape) flipRotations() []Shape {
	transforms := []Shape{shape}
	flip := make(Shape, len(shape))
	copy(flip, shape)
	for i, point := range shape {
		flip[i].X = point.Y
		flip[i].Y = point.X
	}
	transforms = append(transforms, flip)
	return transforms
}

func (shape Shape) identityRotations() []Shape {
	return []Shape{shape}
}
