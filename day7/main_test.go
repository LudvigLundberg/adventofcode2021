package main

import "testing"

func TestCalculateFuelDistance(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	expect := 37

	got := calculateFuelDistance(input, distanceFromEach)

	if expect != got {
		t.Errorf("expect: %v, got: %v", expect, got)
	}
}

func TestCalculateFuelDistance2(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	expect := 168

	got := calculateFuelDistance(input, distanceFromEach2)

	if expect != got {
		t.Errorf("expect: %v, got: %v", expect, got)
	}
}
