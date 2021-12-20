package main

import (
	"fmt"
	"sort"

	"github.com/LudvigLundberg/adventofcode2021/parseinput"
)

func main() {
	input, err := parseinput.ParseFile("input")
	if err != nil {
		panic(err)
	}

	partOne(input)
	partTwo(input)
}

func partOne(input []string) {
	fmt.Println("-------- Part One -------")

	sum := 0
	for _, line := range input {
		value, _ := validate(line)
		sum += value
	}

	fmt.Printf("Sum of illegals: %v\n", sum)
}

func partTwo(input []string) {
	fmt.Println("-------- Part Two -------")

	values := make([]int, 0, len(input))
	for _, line := range input {
		_, unfinished := validate(line)
		if len(unfinished) != 0 {
			values = append(values, complete(unfinished))
		}
	}
	sort.Ints(values)
	fmt.Printf("Middle value of unfinished: %v\n", values[(len(values)/2)])
}

const (
	LEFT_PAREN    = '('
	RIGHT_PAREN   = ')'
	LEFT_BRACKET  = '['
	RIGHT_BRACKET = ']'
	LEFT_CURLY    = '{'
	RIGHT_CURLY   = '}'
	LESS_THAN     = '<'
	GREATER_THAN  = '>'
)

func validate(line string) (int, []rune) {
	current := 0
	expectedTokens := make([]rune, len(line))

	for _, r := range line {
		switch r {
		case LEFT_PAREN:
			expectedTokens[current] = r + 1
			current++
		case LEFT_BRACKET, LEFT_CURLY, LESS_THAN:
			expectedTokens[current] = r + 2
			current++

		case RIGHT_PAREN, RIGHT_BRACKET, RIGHT_CURLY, GREATER_THAN:
			if expectedTokens[current-1] == r {
				current--
			} else {
				return charValue(r), []rune{}
			}
		}
	}
	if current == 0 {
		return 0, []rune{}
	} else {
		return 0, reverse(expectedTokens[:current])
	}
}

func reverse(input []rune) []rune {
	reversed := make([]rune, len(input))
	last := len(input) - 1
	for i := range reversed {
		reversed[i] = input[last-i]
	}
	return reversed
}

func complete(tokens []rune) int {
	sum := 0
	for _, r := range tokens {
		sum = sum * 5
		switch r {
		case RIGHT_PAREN:
			sum += 1
		case RIGHT_BRACKET:
			sum += 2
		case RIGHT_CURLY:
			sum += 3
		case GREATER_THAN:
			sum += 4
		}
	}
	return sum
}

func charValue(token rune) int {
	switch token {
	case RIGHT_PAREN:
		return 3
	case RIGHT_BRACKET:
		return 57
	case RIGHT_CURLY:
		return 1197
	case GREATER_THAN:
		return 25137
	default:
		panic("Invalid character")
	}
}
