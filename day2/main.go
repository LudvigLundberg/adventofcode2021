package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/LudvigLundberg/adventofcode2021/parseinput"
)

type Direction = string

const (
	Forward Direction = "forward"
	Up      Direction = "up"
	Down    Direction = "down"
)

type Movement struct {
	Direction Direction
	Value     int
}

func getCommands(input []string) ([]Movement, error) {
	commands := make([]Movement, len(input))

	for i, line := range input {
		split := strings.Split(line, " ")
		value, err := strconv.Atoi(split[1])

		if err != nil {
			return nil, err
		}
		commands[i] = Movement{Direction: split[0], Value: value}
	}
	return commands, nil
}

func partOne(input []Movement) {
	var horizontal, vertical int

	for _, movement := range input {
		switch movement.Direction {
		case Forward:
			horizontal += movement.Value
		case Up:
			vertical -= movement.Value
		case Down:
			vertical += movement.Value
		}
	}

	fmt.Printf("horizontal: %d, vertical: %d, multiplied: %d\n", horizontal, vertical, horizontal*vertical)
}

func partTwo(input []Movement) {
	var horizontal, vertical, aim int

	for _, movement := range input {
		switch movement.Direction {
		case Forward:
			horizontal += movement.Value
			vertical += aim * movement.Value
		case Up:
			aim -= movement.Value
		case Down:
			aim += movement.Value
		}
	}
	fmt.Printf("horizontal: %d, vertical: %d, aim: %d, multiplied: %d\n", horizontal, vertical, aim, horizontal*vertical)
}

func main() {
	input, err := parseinput.ParseFile("input")

	if err != nil {
		log.Fatalf("error getting input: %v", err)
	}
	commands, err := getCommands(input)

	if err != nil {
		log.Fatalf("error getting commands: %v", err)
	}
	fmt.Println("---------- Part One ---------")
	partOne(commands)

	fmt.Println("---------- Part Two ---------")
	partTwo(commands)
}
