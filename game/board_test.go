package game

import (
	"reflect"
	"testing"
)

func TestFilledRowIndices(t *testing.T) {
	board := &Board{}
	board.Reset()

	rows := []int{0, 2, 4}

	// fill a line with ones
	for i := 0; i < BoardWidth; i++ {
		for _, v := range rows {
			board.Grid[v][i] = 1
		}
	}

	result := board.filledRowIndices()
	if !reflect.DeepEqual(rows, result) {
		t.Errorf("expected rows %v, found %v", rows, result)
	}
}

func TestClearFilledRows(t *testing.T) {
	board := &Board{}
	board.Reset()

	for i := 0; i < BoardWidth; i++ {
		board.Grid[0][i] = 1
	}

	board.clearFilledRows()
	for i := 0; i < BoardWidth; i++ {
		if board.Grid[0][i] != 0 {
			t.Errorf("failed to clear row")
		}
	}

	if board.ClearedRows != 1 {
		t.Errorf("did not update Board.ClearedRows")
	}
}
