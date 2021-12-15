package main

import (
	"reflect"
	"testing"
)

func TestNumberOfLowPoints(t *testing.T) {
	input := []int{2, 1, 9, 9, 9, 4, 3, 2, 1, 0}

	expect := []int{1, 0}
	got := getLowPoints(input, 10)

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expect: %v, got: %v", expect, got)
	}
}

func TestNumberOfLowPoints2(t *testing.T) {
	input := []int{
		2, 1, 9, 9, 9, 4, 3, 2, 1, 0,
		3, 9, 8, 7, 8, 9, 4, 9, 2, 1,
		9, 8, 5, 6, 7, 8, 9, 8, 9, 2,
		8, 7, 6, 7, 8, 9, 6, 7, 8, 9,
		9, 8, 9, 9, 9, 6, 5, 6, 7, 8,
	}

	expect := []int{1, 0, 5, 5}
	got := getLowPoints(input, 10)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect: %v, got: %v", expect, got)
	}
}
