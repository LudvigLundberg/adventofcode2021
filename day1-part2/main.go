package main

import (
	"fmt"

	"github.com/LudvigLundberg/adventofcode2021/parseinput"
)

func sum(numbers []int) int {
	var sum int
	for _, nr := range numbers {
		sum += nr
	}

	return sum
}

func main() {
	numbers, err := parseinput.ParseFileAsInt("input")

	if err != nil {
		panic("unable to read input")
	}

	var currentsum, increases int
	if len(numbers) > 2 {
		currentsum = sum(numbers[0:3])
		for i := 1; i+2 < len(numbers); i++ {
			numberSum := sum(numbers[i : i+3])
			if numberSum > currentsum {
				increases++
			}
			currentsum = numberSum
		}
	}

	fmt.Printf("Number of increases %d\n", increases)
}
