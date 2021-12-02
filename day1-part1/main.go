package main

import (
	"fmt"
	"strconv"

	"github.com/LudvigLundberg/adventofcode2021/parseinput"
)

func main() {
	input, _ := parseinput.ParseFile("input")

	var greater int

	if len(input) > 0 {
		previous, err := strconv.Atoi(input[0])
		if err != nil {
			panic("cannot convert input")
		}
		for _, line := range input[1:] {
			next, err := strconv.Atoi(line)

			if err != nil {
				panic("cannot convert input")
			}

			if next > previous {
				greater++
			}

			previous = next
		}
	} else {
		panic("input of size 0")
	}

	fmt.Printf("nr of increased consecutive numbers: %v\n", greater)
}
