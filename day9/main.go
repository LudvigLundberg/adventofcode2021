package main

import (
	"fmt"
	"sort"
	"sync"

	"github.com/LudvigLundberg/adventofcode2021/parseinput"
)

func main() {
	input, err := parseinput.ParseFile("input")
	if err != nil {
		panic(err)
	}

	solve(input)

}

func solve(input []string) {
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
		sum += numbers[point] + 1
	}
	fmt.Println("-------- Part One ---------")
	fmt.Printf("Points %v, (element + 1) sum: %v\n", len(points), sum)

	fmt.Println("-------- Part Two ---------")
	basins := getBasins(numbers, points, rowLen)
	fmt.Printf("Basin sizes: %v, multiplied: %v\n", basins, multiply(basins))
}

func getLowPoints(input []int, rowSize int) []int {
	lowPoints := make([]int, 0, 100)

	for i, nr := range input {
		compareToIdexes := neighbors(i, input, rowSize)
		compareTo := mapTo(compareToIdexes, func(i int) int {
			return input[i]
		})
		if lowerThanAll(nr, compareTo...) {
			lowPoints = append(lowPoints, i)
		}
	}

	return lowPoints
}

func mapTo(xs []int, fn func(int) int) []int {
	ys := make([]int, len(xs))
	for i, x := range xs {
		ys[i] = fn(x)
	}
	return ys
}

func reduce(xs []int, fn func(int, int) int) int {
	y := xs[0]

	for _, x := range xs[1:] {
		y = fn(y, x)
	}
	return y
}

func multiply(xs []int) int {
	return reduce(xs, func(x, y int) int { return x * y })
}

func getBasins(input []int, lowPoints []int, rowSize int) []int {
	wg := sync.WaitGroup{}
	wg.Add(len(lowPoints))

	sizes := make([]int, len(lowPoints))

	for i, point := range lowPoints {
		go func(point int, index int) {
			sizes[index] = explore(point, input, rowSize)
			wg.Done()
		}(point, i)
	}
	wg.Wait()

	sort.Ints(sizes)

	return sizes[len(sizes)-3:]
}

func explore(start int, input []int, rowSize int) int {
	explored := make(map[int]bool)

	stack := make([]int, 1, 100)
	stack[0] = start
	explored[start] = true
	size := 1
	for last := 0; last > -1; last = len(stack) - 1 {
		current := stack[last]
		stack = stack[:last]
		neighbors := neighbors(current, input, rowSize)
		for _, neighbor := range neighbors {
			if input[neighbor] != 9 && !explored[neighbor] {
				explored[neighbor] = true
				stack = append(stack, neighbor)
				size++
			}
		}
	}

	return size
}

func neighbors(i int, input []int, rowSize int) []int {
	neighbors := make([]int, 0, 4)
	if (i%rowSize)-1 >= 0 {
		neighbors = append(neighbors, i-1)
	}

	if (i%rowSize)+1 < rowSize {
		neighbors = append(neighbors, i+1)
	}

	if i-rowSize >= 0 {
		neighbors = append(neighbors, i-rowSize)
	}

	if i+rowSize < len(input) {
		neighbors = append(neighbors, i+rowSize)
	}
	return neighbors
}

func lowerThanAll(nr int, numbers ...int) bool {
	for _, compare := range numbers {
		if nr >= compare {
			return false
		}
	}
	return true
}
