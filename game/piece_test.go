package game

import (
	"reflect"
	"testing"
)

func TestIdentityRotations(t *testing.T) {
	input := Shape{
		{2, 4},
		{6, 8},
		{10, 12},
	}

	expected := []Shape{input}
	result := input.identityRotations()

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestFlipRotations(t *testing.T) {
	input := Shape{
		{0, 1},
		{2, 3},
	}

	expected := []Shape{
		input,
		{
			{1, 0},
			{3, 2},
		},
	}

	result := input.flipRotations()

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestSquareRotations(t *testing.T) {
	input := Shape{{1, 0}}
	expected := []Shape{
		input,
		{{2, 1}},
		{{1, 2}},
		{{0, 1}},
	}

	result := input.squareRotations()
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
