package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/LudvigLundberg/adventofcode2021/parseinput"
)

func main() {
	input, err := parseinput.ParseFile("input")

	if err != nil {
		panic(err)
	}

	fmt.Println("------- Part One ------")
	partOne(input)

	fmt.Println("------- Part Two -------")

	partTwo(input)
}

func partOne(input []string) {
	input = strings.Split(input[0], ",")
	numbers := make([]int, len(input))

	for i, line := range input {
		number, err := strconv.Atoi(line)

		if err != nil {
			panic(err)
		}

		numbers[i] = number
	}

	shortestDistance := calculateFuelDistance(numbers, distanceFromEach)

	fmt.Printf("shortest distance of input: %v\n", shortestDistance)
}

func partTwo(input []string) {
	input = strings.Split(input[0], ",")
	numbers := make([]int, len(input))

	for i, line := range input {
		number, err := strconv.Atoi(line)

		if err != nil {
			panic(err)
		}

		numbers[i] = number
	}

	shortestDistance := calculateFuelDistance(numbers, distanceFromEach2)

	fmt.Printf("shortest distance of input: %v\n", shortestDistance)

}

func calculateFuelDistance(input []int, distanceToElement func([]int, int) int) int {
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)

	shortest := MaxInt
	for i := range input {
		distance := distanceToElement(input, i)
		if distance < shortest {
			shortest = distance
		}
	}
	return shortest
}

func distanceFromEach(elements []int, to int) int {
	var totalDistance int
	for _, element := range elements {
		totalDistance += abs(to - element)
	}

	return totalDistance
}

func distanceFromEach2(elements []int, to int) int {
	var totalDistance int
	for _, element := range elements {
		totalDistance += sum(abs(to - element))
	}
	return totalDistance
}

func abs(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}

func sum(nr int) int {
	if nr == 1 || nr == 0 {
		return 1
	} else {
		return sum(nr-1) + nr
	}
}
