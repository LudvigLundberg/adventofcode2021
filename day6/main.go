package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/LudvigLundberg/adventofcode2021/parseinput"
)

func main() {
	stringInput, err := parseinput.ParseFile("input")
	if err != nil {
		panic(err)
	}

	input := make([]int, 0, 500)
	numbers := strings.Split(stringInput[0], ",")
	for _, number := range numbers {
		number, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		input = append(input, number)
	}
	fmt.Println("-------- Part One ------")
	partOne(input, 80)

	fmt.Println("-------- Part Two ------")
	partOne(input, 256)
}

func partOne(input []int, simulatedays int) {
	counts := make([]int, 9)
	for _, nr := range input {
		counts[nr] = counts[nr] + 1
	}

	for i := 0; i < simulatedays; i++ {
		counts = SimulateDay(counts)
	}

	totalCount := 0
	for _, count := range counts {
		totalCount += count
	}
	fmt.Printf("total number of fish: %v\n", totalCount)
}

func SimulateDay(currentDay []int) []int {
	newDay := make([]int, len(currentDay))
	copy(newDay, currentDay[1:])
	newDay[len(newDay)-1] = currentDay[0]
	newDay[len(newDay)-3] = currentDay[len(currentDay)-2] + currentDay[0]
	return newDay
}
