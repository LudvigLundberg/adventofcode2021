package main

import (
	"reflect"
	"testing"
)

func TestNumberOfFlashes(t *testing.T) {
	input := [][]int{
		{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
		{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
		{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
		{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
		{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
		{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
		{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
		{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
		{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
		{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
	}

	expect := 1656
	var flashes int
	totalFlashes := 0
	for i := 1; i <= 100; i++ {
		input, flashes, _ = flash(input)
		totalFlashes += flashes
	}

	if expect != totalFlashes {
		t.Errorf("expect: %v, got: %v", expect, totalFlashes)
	}
}

func TestStep(t *testing.T) {
	input := [][]int{{1, 1, 1, 1, 1},
		{1, 9, 9, 9, 1},
		{1, 9, 1, 9, 1},
		{1, 9, 9, 9, 1},
		{1, 1, 1, 1, 1},
	}

	expect := [][]int{
		{3, 4, 5, 4, 3},
		{4, 0, 0, 0, 4},
		{5, 0, 0, 0, 5},
		{4, 0, 0, 0, 4},
		{3, 4, 5, 4, 3},
	}

	expectFlashes := 9

	got, flashes, _ := flash(input)

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expect: %v, got %v", expect, got)
	}

	if flashes != expectFlashes {
		t.Errorf("expected: %v flashes, got: %v flashes", expectFlashes, flashes)
	}
}

func TestAllFlashed(t *testing.T) {
	input := [][]int{
		{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
		{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
		{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
		{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
		{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
		{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
		{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
		{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
		{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
		{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
	}
	var allFlashed bool

	for i := 1; i <= 195; i++ {
		input, _, allFlashed = flash(input)
	}

	if !allFlashed {
		t.Errorf("expected all to flash: %v", input)
	}
}
