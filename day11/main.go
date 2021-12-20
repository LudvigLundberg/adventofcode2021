package main

import (
	"fmt"

	"github.com/LudvigLundberg/adventofcode2021/parseinput"
)

func main() {
	input, err := parseinput.ParseFile("input")

	if err != nil {
		panic(err)
	}

	matrix := make(Matrix, len(input))
	for i, line := range input {
		row := make([]int, len(line))
		for j, r := range line {
			row[j] = int(r - '0')
		}
		matrix[i] = row
	}

	partOne(matrix)
	partTwo(matrix)
}

func partOne(matrix Matrix) {
	fmt.Println("------------ Part One -------------")

	var totalFlashes, flashes int

	for i := 1; i <= 100; i++ {
		matrix, flashes, _ = flash(matrix)
		totalFlashes += flashes
	}

	fmt.Printf("total number of flashes: %v\n", totalFlashes)
}

func partTwo(matrix Matrix) {
	fmt.Println("------------ Part Two -------------")

	var allFlashed bool

	for i := 1; ; i++ {
		matrix, _, allFlashed = flash(matrix)
		if allFlashed {
			fmt.Printf("All flas triggered on loop: %v\n", i)
			return
		}
	}

}

type Position struct {
	x, y int
}

type Matrix = [][]int

func flash(in Matrix) (Matrix, int, bool) {
	input := make(Matrix, len(in))
	for i, row := range in {
		input[i] = make([]int, len(row))
		copy(input[i], row)
	}

	triggered := make([]Position, 0, len(input)*100)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			input[i][j]++
			if input[i][j] == 10 {
				triggered = append(triggered, neighbors(Position{i, j}, len(input))...)
			}
		}
	}
	if len(triggered) != 0 {

		for x := 0; x < len(triggered); x++ {
			i, j := triggered[x].x, triggered[x].y
			switch input[i][j] {
			case 10:
			case 9:
				input[i][j]++
				triggered = append(triggered, neighbors(Position{i, j}, len(input))...)
			default:
				input[i][j]++
			}
		}
	}

	flashes := 0
	allFlashed := true
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 10 {
				input[i][j] = 0
				flashes++
			} else {
				allFlashed = false
			}
		}
	}
	return input, flashes, allFlashed
}

func neighbors(p Position, length int) []Position {
	output := make([]Position, 0, 8)

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {

			switch pos := (Position{p.x + i, p.y + j}); {
			case pos == p, pos.x < 0, pos.x >= length, pos.y < 0, pos.y >= length:
				continue
			default:
				output = append(output, pos)
			}
		}
	}
	return output
}
