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

	foldStrings, err := parseinput.ParseFile("folds")

	if err != nil {
		panic(err)
	}

	paper := createPaper(input)
	folds := createFolds(foldStrings)

	partOne(paper, folds)
	partTwo(paper, folds)
}

func partOne(paper TransparentPaper, folds []Fold) {
	fmt.Println("------------- Part One -----------")

	fold(paper, folds[0])

	filled := len(paper)

	fmt.Printf("Nr of elements after 1 fold: %v\n", filled)
}

func partTwo(paper TransparentPaper, folds []Fold) {
	fmt.Println("------------ Part Two ----------")

	for _, AFold := range folds[1:] {
		fold(paper, AFold)
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			_, exist := paper[Dot{j, i}]

			if exist {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println()
	}
}

type TransparentPaper = map[Dot]struct{}
type FoldType = int
type Dot struct {
	x, y int
}

const (
	Horizontal = iota
	Vertical
)

type Fold struct {
	line   int
	typeOf FoldType
}

func createFolds(lines []string) []Fold {
	folds := make([]Fold, len(lines))

	for i, line := range lines {
		var fold Fold
		split := strings.Split(line, "=")
		direction := split[0]
		nr, err := strconv.Atoi(split[1])

		if err != nil {
			panic(err)
		}

		switch direction {
		case "y":
			fold = Fold{line: nr, typeOf: Horizontal}
		case "x":
			fold = Fold{line: nr, typeOf: Vertical}
		}

		folds[i] = fold
	}

	return folds
}

func createPaper(lines []string) TransparentPaper {

	paper := make(TransparentPaper, len(lines))

	for _, line := range lines {
		split := strings.Split(line, ",")
		x, err := strconv.Atoi(split[0])

		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(split[1])

		if err != nil {
			panic(err)
		}
		paper[Dot{x, y}] = struct{}{}
	}

	return paper
}

func fold(paper TransparentPaper, fold Fold) {

	for dot := range paper {
		switch fold.typeOf {
		case Vertical:
			if dot.x > fold.line {
				delete(paper, dot)
				newX := fold.line - (dot.x - fold.line)
				paper[Dot{newX, dot.y}] = struct{}{}
			}
		case Horizontal:
			if dot.y > fold.line {
				delete(paper, dot)
				newY := fold.line - (dot.y - fold.line)
				paper[Dot{dot.x, newY}] = struct{}{}
			}
		}
	}
}
