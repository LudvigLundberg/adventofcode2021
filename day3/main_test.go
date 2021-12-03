package main

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	verify := []byte{'1', '0', '1', '0', '1'}

	input := make(chan byte, 5)
	returnChan := make(chan byte, 1)

	go countBits(input, returnChan)

	for _, r := range verify {
		input <- r
	}
	close(input)

	var expect byte
	expect = '1'
	got := <-returnChan

	if expect != got {
		t.Errorf("expected: %s, got: %s", string(expect), string(got))
	}
}

func TestFilterByFirstBit(t *testing.T) {
	verify := []string{"1001", "0101", "1010"}
	expect := []string{"1001", "1010"}

	got := filterByBit(verify, '1', 0)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expected: %v, got: %v", expect, got)
	}
}
