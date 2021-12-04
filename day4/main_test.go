package main

import (
	"reflect"
	"testing"
)

func createTestBingo() bingo {
	return bingo{
		board: [5][5]bool{},
		numbers: map[int]position{
			22: {0, 0},
			13: {0, 1},
			17: {0, 2},
			11: {0, 3},
			0:  {0, 4},
			8:  {1, 0},
			2:  {1, 1},
			23: {1, 2},
			4:  {1, 3},
			24: {1, 4},
			21: {2, 0},
			9:  {2, 1},
			14: {2, 2},
			16: {2, 3},
			7:  {2, 4},
			6:  {3, 0},
			10: {3, 1},
			3:  {3, 2},
			18: {3, 3},
			5:  {3, 4},
			1:  {4, 0},
			12: {4, 1},
			20: {4, 2},
			15: {4, 3},
			19: {4, 4},
		},
	}
}

func TestPartOne(t *testing.T) {

}

func TestPartTwo(t *testing.T) {

}

func TestFill(t *testing.T) {
	bingo := createTestBingo()

	bingo.Fill(22)

	if bingo.board[0][0] != true {
		t.Errorf("expect %v for value %v, got: %v", true, 22, false)
	}

	bingo.Fill(14)

	if bingo.board[2][2] != true {
		t.Errorf("expect %v for value %v, got: %v", true, 14, false)
	}

	bingo.Fill(19)

	if bingo.board[4][4] != true {
		t.Errorf("expect %v for value %v, got: %v", true, 19, false)
	}
}

func TestFillError(t *testing.T) {
	bingo := createTestBingo()

	_, err := bingo.Fill(100)

	if err == nil {
		t.Errorf("expected error when number not on board")
	}
}

func TestCheckWin(t *testing.T) {
	bingo := createTestBingo()

	bingo.Fill(22)
	bingo.Fill(13)
	bingo.Fill(17)
	bingo.Fill(11)
	win, _ := bingo.Fill(0)

	if !win {
		t.Errorf("expected win: %v, got win: %v", true, win)
	}
}

func TestCreateBingo(t *testing.T) {
	input := [5][5]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19},
	}
	got := CreateBingo(input)
	expect := createTestBingo()

	if reflect.DeepEqual(expect, got) {
		t.Errorf("expected: %v, got %v", expect, got)
	}
}
