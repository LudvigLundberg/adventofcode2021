package main

import (
	"reflect"
	"testing"
)

func TestNumberOfPaths(t *testing.T) {
	input := map[string][]string{
		"start": {"A", "b"},
		"A":     {"start", "c", "b", "end"},
		"b":     {"start", "A", "d", "end"},
		"d":     {"b"},
		"c":     {"A"},
	}

	expect := 10

	got := numberOfPaths(input, "start", "end", map[string]bool{}, true)

	if expect != got {
		t.Errorf("expect: %v, got: %v", expect, got)
	}
}

func TestCreateGraph(t *testing.T) {
	input := []Edge{
		{"start", "A"},
		{"start", "b"},
		{"A", "c"},
		{"A", "b"},
		{"b", "d"},
		{"A", "end"},
		{"b", "end"},
	}
	expect := map[string][]string{
		"start": {"A", "b"},
		"A":     {"start", "c", "b", "end"},
		"b":     {"start", "A", "d", "end"},
		"d":     {"b"},
		"c":     {"A"},
		"end":   {"A", "b"},
	}

	got := createGraph(input)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect: %v, got %v", expect, got)
	}
}
