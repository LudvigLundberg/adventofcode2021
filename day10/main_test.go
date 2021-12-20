package main

import (
	"reflect"
	"testing"
)

func TestValidLine(t *testing.T) {
	input := "[<>({}){}[([])<>]]"

	expect := 0
	got, _ := validate(input)

	if expect != got {
		t.Errorf("expect: %v, got: %v", expect, got)
	}
}

func TestInvalidLine(t *testing.T) {
	input := "<)"

	expect := 3
	got, _ := validate(input)

	if expect != got {
		t.Errorf("expect: %v, got: %v", expect, got)
	}
}

func TestNonCompleteLine(t *testing.T) {
	input := "(((("

	expect := []rune{')', ')', ')', ')'}

	_, got := validate(input)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect: %v, got: %v", expect, got)
	}

	input = "(((({<>}<{<{<>}{[]{[]{}"

	expect = []rune{'}', '}', '>', '}', '>', ')', ')', ')', ')'}

	_, got = validate(input)

	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect: %v, got: %v", expect, got)
	}
}

func TestCalculateRemainder(t *testing.T) {
	input := []rune{']', ')', '}', '>'}

	expect := 294
	got := complete(input)

	if expect != got {
		t.Errorf("expect: %v, got: %v", expect, got)
	}

	input = []rune{'}', '}', ']', ']', ')', '}', ')', ']'}

	expect = 288957
	got = complete(input)

	if expect != got {
		t.Errorf("expect: %v, got: %v", expect, got)
	}

	input = []rune{'}', '}', '>', '}', '>', ')', ')', ')', ')'}

	expect = 1480781
	got = complete(input)

	if expect != got {
		t.Errorf("expect: %v, got: %v", expect, got)
	}

	input = []rune{']', ']', '}', '}', ']', '}', ']', '}', '>'}

	expect = 995444
	got = complete(input)

	if expect != got {
		t.Errorf("expect: %v, got: %v", expect, got)
	}
}
