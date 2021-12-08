package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/LudvigLundberg/adventofcode2021/parseinput"
)

func main() {
	input := parseinput.ParseFile("input")
}

func partOne(input []string) {

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
	case start.x < end.x && start.y < end.y:
		return diagonal(start, end), nil
	case start.x < end.x:
		return horizontal(start, end), nil
	case start.y < end.y:
		return vertical(start, end), nil
	}
	return nil, fmt.Errorf("Invalid state of line, start: %v, end: %v", start, end)
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

func diagonal(start, end Point) []Point {
	return []Point{}
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

	for _, line := range input {
		points := strings.Split(line, " -> ")
		x, err := extractPoint(points[0])

		if err != nil {
			return []Line{}, err
		}
		y, err := extractPoint(points[1])

		if err != nil {
			return []Line{}, err
		}

		lines = append(lines, Line{x, y})
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

	return Point{x, y}, nil
}
