package game

import (
	"math/rand"
	"time"
)

const (
	BoardHeight = 20
	BoardWidth  = 10
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type Board struct {
	Grid        [][]int
	Piece       Piece
	ClearedRows int
}

// Reset clears the board and selects the initial piece.
func (board *Board) Reset() {
	board.ClearedRows = 0
	board.Grid = make([][]int, BoardHeight)
	for i := 0; i < BoardHeight; i++ {
		board.Grid[i] = make([]int, BoardWidth)
	}
	board.NextPiece()
}

// NextPiece randomly selects a new piece.
func (board *Board) NextPiece() bool {
	blockValue := int(rand.Intn(256))
	rotations := pieceRotations[rand.Intn(len(pieceRotations))]
	coords := Point{(BoardWidth / 2) - 1, 0}

	board.Piece = Piece{blockValue, coords, 0, rotations}
	if board.Piece.overlapsGrid(board.Grid) {
		return false
	}
	return true
}

// Level is the current game level, based on cleared rows.
func (board *Board) Level() int {
	return board.ClearedRows / 10
}

// clearFilledRows removes filled rows from the Board.
func (board *Board) clearFilledRows() {
	rows := board.filledRowIndices()
	if len(rows) == 0 {
		return
	}

	for _, v := range rows {
		board.clearRow(v)
	}
}

// DropPiece lowers the current piece by one square. If the piece cannot
// move any lower, it adds the piece to the grid and clears out filled
// rows on the board.
func (board *Board) DropPiece() bool {
	board.Piece.coords.Y++
	if !board.Piece.overlapsGrid(board.Grid) {
		return true
	}

	board.Piece.coords.Y--
	shape := board.Piece.GetShape()

	for _, point := range shape {
		board.Grid[point.Y][point.X] = board.Piece.Value
	}

	board.clearFilledRows()
	return false
}

// QuickDropPiece drops a piece until it lands in its final resting place.
func (board *Board) QuickDropPiece() {
	for board.DropPiece() {
	}
}

func (board *Board) MovePieceLeft() bool {
	board.Piece.coords.X--
	if board.Piece.overlapsGrid(board.Grid) {
		board.Piece.coords.X++
		return false
	}
	return true
}

func (board *Board) MovePieceRight() bool {
	board.Piece.coords.X++
	if board.Piece.overlapsGrid(board.Grid) {
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

	if board.Piece.overlapsGrid(board.Grid) {
		board.Piece.rotationIndex++
		return false
	}
	return true
}

func (board *Board) RotatePieceRight() bool {
	board.Piece.rotationIndex++
	if board.Piece.overlapsGrid(board.Grid) {
		board.Piece.rotationIndex--
		return false
	}
	return true
}

func (board *Board) filledRowIndices() []int {
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
	board.Grid[0] = make([]int, BoardWidth)
	board.ClearedRows++
	return true
}
