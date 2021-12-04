package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/LudvigLundberg/adventofcode2021/parseinput"
)

func main() {
	fmt.Println("---------- Part One ---------")
	partOne()
	fmt.Println("---------- Part Two ---------")
	partTwo()
}

func partOne() {
	input, err := parseinput.ParseFile("input")
	if err != nil {
		panic(err)
	}
	drawnNumbers, err := generateNumbers(input[0])

	if err != nil {
		panic(err)
	}

	bingos := make([]Bingo, 0, 20)

	for i := 1; i < len(input); i += 6 {
		board, err := generateBoard(input[i+1 : i+6])
		if err != nil {
			panic(err)
		}
		bingos = append(bingos, CreateBingo(board))
	}

	for _, nr := range drawnNumbers {
		for _, bingo := range bingos {
			win, _ := bingo.Fill(nr)
			if win {
				score := bingo.Score(nr)
				fmt.Printf("total score: %v\n", score)
				return
			}
		}
	}
}

func partTwo() {
	input, err := parseinput.ParseFile("input")
	if err != nil {
		panic(err)
	}
	drawnNumbers, err := generateNumbers(input[0])

	if err != nil {
		panic(err)
	}

	bingos := make([]Bingo, 0, 20)
	winners := make(map[Bingo]struct{})
	lastScore := 0

	for i := 1; i < len(input); i += 6 {
		board, err := generateBoard(input[i+1 : i+6])
		if err != nil {
			panic(err)
		}
		bingos = append(bingos, CreateBingo(board))
	}

	for _, nr := range drawnNumbers {
		for _, bingo := range bingos {
			win, _ := bingo.Fill(nr)
			if win {
				_, exists := winners[bingo]
				if !exists {
					lastScore = bingo.Score(nr)
					winners[bingo] = struct{}{}
				}
			}
		}
	}

	fmt.Printf("total score of last winner: %v\n", lastScore)
}

func generateBoard(input []string) ([5][5]int, error) {
	numbers := [5][5]int{}
	for i, line := range input {
		split := strings.Fields(line)

		for j, nrString := range split {
			nr, err := strconv.Atoi(nrString)
			if err != nil {
				return [5][5]int{}, err
			}
			numbers[i][j] = nr
		}
	}

	return numbers, nil
}

func generateNumbers(line string) ([]int, error) {
	split := strings.Split(line, ",")
	numbers := make([]int, len(split))

	for i := range split {
		nr, err := strconv.Atoi(split[i])
		if err != nil {
			return []int{}, err
		}
		numbers[i] = nr
	}
	return numbers, nil
}

func CreateBingo(board [5][5]int) Bingo {
	bingo := bingo{board: [5][5]bool{}, numbers: make(map[int]position)}
	for rowIndex, row := range board {
		for columnIndex, nr := range row {
			bingo.numbers[nr] = position{row: rowIndex, column: columnIndex}
		}
	}
	return &bingo
}

type Bingo interface {
	Fill(int) (bool, error)
	Score(int) int
}

type board = [5][5]bool

type position struct {
	row    int
	column int
}

type bingo struct {
	board   board
	numbers map[int]position
}

func (b *bingo) Fill(drawn int) (bool, error) {
	position, ok := b.numbers[drawn]

	if !ok {
		return false, fmt.Errorf("number does not exist on board")
	}

	b.board[position.row][position.column] = true

	return b.checkWin(position), nil
}

func (b *bingo) checkWin(p position) bool {
	return b.rowWin(p.row) || b.columnWin(p.column)
}

func (b *bingo) rowWin(row int) bool {
	for _, filled := range b.board[row] {
		if !filled {
			return false
		}
	}
	return true
}

func (b *bingo) columnWin(column int) bool {
	for _, row := range b.board {
		if !row[column] {
			return false
		}
	}
	return true
}

func (b *bingo) Score(drawn int) int {
	sum := 0
	for nr, position := range b.numbers {
		if !b.board[position.row][position.column] {
			sum += nr
		}
	}
	return sum * drawn
}
