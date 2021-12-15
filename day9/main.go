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
	fmt.Println("-------- Part One ---------")
	partOne(input)
}

func partOne(input []string) {
	rowLen := len(input[0])

	numbers := make([]int, 0, rowLen*100)

	for _, line := range input {
		for _, r := range line {
			numbers = append(numbers, int(r-'0'))
		}
	}

	points := getLowPoints(numbers, rowLen)

	sum := 0

	for _, point := range points {
		sum += point + 1
	}

	fmt.Printf("Points %v, (element + 1) sum: %v\n", len(points), sum)
}

func getLowPoints(input []int, rowSize int) []int {
	lowPoints := make([]int, 0, 100)

	for i, nr := range input {
		compareTo := make([]int, 0, 4)
		if (i%rowSize)-1 >= 0 {
			compareTo = append(compareTo, input[i-1])
		}

		if (i%rowSize)+1 < rowSize {
			compareTo = append(compareTo, input[i+1])
		}

		if i-rowSize >= 0 {
			compareTo = append(compareTo, input[i-rowSize])
		}

		if i+rowSize < len(input) {
			compareTo = append(compareTo, input[i+rowSize])
		}

		if lowerThanAll(nr, compareTo...) {
			lowPoints = append(lowPoints, nr)
		}
	}

	return lowPoints
}

func lowerThanAll(nr int, numbers ...int) bool {
	for _, compare := range numbers {
		if nr >= compare {
			return false
		}
	}
	return true
}
