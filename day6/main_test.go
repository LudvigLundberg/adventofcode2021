package main

import (
	"reflect"
	"testing"
)

func TestSimulateDay(t *testing.T) {
	start := []int{0, 1, 1, 2, 1, 0, 0, 0, 0}
	expect := []int{1, 1, 2, 1, 0, 0, 0, 0, 0}

	got := SimulateDay(start)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect: %v, got: %v", expect, got)
	}
}

func TestSimulateDayWithExpansion(t *testing.T) {
	start := []int{1, 1, 2, 1, 0, 0, 0, 0, 0}
	expect := []int{1, 2, 1, 0, 0, 0, 1, 0, 1}

	got := SimulateDay(start)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect: %v, got: %v", expect, got)
	}
}

func TestSimulate80DaysCount(t *testing.T) {
	start := []int{0, 1, 1, 2, 1, 0, 0, 0, 0}

	for i := 0; i < 80; i++ {
		start = SimulateDay(start)
	}
	expect := 5934
	var got int

	for _, count := range start {
		got += count
	}

	if got != expect {
		t.Errorf("expect: %v, got: %v", expect, got)
	}
}
