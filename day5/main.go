package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/LudvigLundberg/adventofcode2021/parseinput"
)

func main() {
	input, err := parseinput.ParseFile("input")
	if err != nil {
		panic(err)
	}

	fmt.Println("------------ Part one --------")
	partOne(input)

	fmt.Println("------------ Part Two --------")
	partTwo(input)
}

func partOne(input []string) {
	lines, err := ToIntervals(input)
	if err != nil {
		panic(err)
	}

	count := make(map[Point]int, len(lines))
	for _, line := range lines {
		points, err := line.AllPoints()
		if err != nil {
			panic(err)
		}
		for _, point := range points {
			if reflect.DeepEqual(point, Point{8, 4}) {
				fmt.Printf("point: %v", point)
			}
			pointCount, ok := count[point]
			if ok {
				count[point] = pointCount + 1
			} else {
				count[point] = 1
			}
		}
	}
	totalcount := 0
	for point, count := range count {

		if count >= 2 {
			fmt.Printf("Point: %v, count: %v\n", point, count)
			totalcount++
		}
	}
	fmt.Printf("Total count: %d\n", totalcount)
}

func partTwo(input []string) {

}

type Point struct {
	x, y int
}
type Line struct {
	from, to Point
}

func (l Line) AllPoints() ([]Point, error) {
	start := l.from
	end := l.to

	switch {
	case start.x != end.x && start.y != end.y:
		return diagonal(start, end), nil
	case start.x != end.x:
		return horizontal(start, end), nil
	case start.y != end.y:
		return vertical(start, end), nil
	}
	return nil, fmt.Errorf("invalid state of line, start: %v, end: %v", start, end)
}

func vertical(start, end Point) []Point {
	start, end = order(start, end)
	points := make([]Point, 0, end.y-start.y+2)
	for i := start.y; i <= end.y; i++ {
		points = append(points, Point{start.x, i})
	}
	return points
}

func order(x, y Point) (Point, Point) {
	if x.x > y.x || x.y > y.y {
		return y, x
	} else {
		return x, y
	}
}

func abs(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}

func diagonal(start, end Point) []Point {
	horizontalIncrement, verticalIncrement := 1, 1
	diff := abs(start.x - end.x)

	points := make([]Point, 0, diff+1)

	if start.x > end.x {
		horizontalIncrement = -1
	}

	if start.y > end.y {
		verticalIncrement = -1
	}

	i := start.x
	j := start.y

	for count := diff + 1; count > 0; count-- {
		points = append(points, Point{i, j})
		i = i + horizontalIncrement
		j = j + verticalIncrement
	}
	return points
}

func horizontal(start, end Point) []Point {
	start, end = order(start, end)
	points := make([]Point, 0, end.x-start.x+2)
	for i := start.x; i <= end.x; i++ {
		points = append(points, Point{i, start.y})
	}
	return points
}

func ToIntervals(input []string) ([]Line, error) {
	lines := make([]Line, len(input))

	for i, line := range input {
		points := strings.Split(line, " -> ")
		x, err := extractPoint(points[0])

		if err != nil {
			return []Line{}, err
		}
		y, err := extractPoint(points[1])

		if err != nil {
			return []Line{}, err
		}

		lines[i] = Line{x, y}
	}

	return lines, nil
}

func extractPoint(s string) (Point, error) {
	coords := strings.Split(s, ",")

	x, err := strconv.Atoi(coords[0])

	if err != nil {
		return Point{}, err
	}

	y, err := strconv.Atoi(coords[1])

	if err != nil {
		return Point{}, err
	}

	return Point{x: x, y: y}, nil
}
