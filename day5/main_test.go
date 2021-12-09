package main

import (
	"reflect"
	"testing"
)

func TestHorizontalPoints(t *testing.T) {
	start := Point{0, 0}
	end := Point{0, 5}

	expect := []Point{
		{0, 0},
		{0, 1},
		{0, 2},
		{0, 3},
		{0, 4},
		{0, 5},
	}

	got := vertical(start, end)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect: %v, got %v", expect, got)
	}
}

func TestVerticalPoints(t *testing.T) {
	start := Point{0, 0}
	end := Point{5, 0}

	expect := []Point{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{5, 0},
	}

	got := horizontal(start, end)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect: %v, got: %v", expect, got)
	}
}

func TestDiagonalPoints(t *testing.T) {
	start := Point{0, 0}
	end := Point{9, 9}

	expect := []Point{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6},
		{7, 7},
		{8, 8},
		{9, 9},
	}

	got := diagonal(start, end)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect: %v, got: %v", expect, got)
	}
}

func TestHorizontalLine(t *testing.T) {
	line := Line{Point{0, 0}, Point{0, 5}}

	expect := []Point{
		{0, 0},
		{0, 1},
		{0, 2},
		{0, 3},
		{0, 4},
		{0, 5},
	}

	got, err := line.AllPoints()

	if err != nil {
		t.Errorf("expected no error: %v", err)
	}
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect: %v, got: %v", expect, got)
	}
}

func TestHorizontalLineReversed(t *testing.T) {
	line := Line{Point{0, 5}, Point{0, 0}}

	expect := []Point{
		{0, 0},
		{0, 1},
		{0, 2},
		{0, 3},
		{0, 4},
		{0, 5},
	}

	got, err := line.AllPoints()

	if err != nil {
		t.Errorf("expected no error: %v", err)
	}
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect: %v, got: %v", expect, got)
	}
}

func TestVerticalLine(t *testing.T) {
	line := Line{Point{0, 0}, Point{5, 0}}

	expect := []Point{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{5, 0},
	}

	got, err := line.AllPoints()

	if err != nil {
		t.Errorf("expected no error: %v", err)
	}
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect: %v, got %v", expect, got)
	}
}

func TestVerticalLineReverse(t *testing.T) {
	line := Line{Point{5, 0}, Point{0, 0}}

	expect := []Point{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{5, 0},
	}

	got, err := line.AllPoints()

	if err != nil {
		t.Errorf("expected no error: %v", err)
	}
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect: %v, got %v", expect, got)
	}
}

func TestDiagonalLine(t *testing.T) {
	line := Line{Point{0, 0}, Point{9, 9}}

	expect := []Point{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6},
		{7, 7},
		{8, 8},
		{9, 9},
	}

	got, err := line.AllPoints()

	if err != nil {
		t.Errorf("expected no error: %v", err)
	}
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect: %v, got: %v", expect, got)
	}
}
func TestDiagonalLineReverse(t *testing.T) {
	line := Line{Point{9, 9}, Point{0, 0}}

	expect := []Point{
		{9, 9},
		{8, 8},
		{7, 7},
		{6, 6},
		{5, 5},
		{4, 4},
		{3, 3},
		{2, 2},
		{1, 1},
		{0, 0},
	}
	got, err := line.AllPoints()

	if err != nil {
		t.Errorf("expected no error: %v", err)
	}
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect: %v, got: %v", expect, got)
	}
}
