package parseinput

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ParseFile attempts to open filepath and returns a slice
// containing each line of the file as a string
func ParseFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	text := make([]string, 0, 500)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text, nil
}

func ParseFileAsInt(filePath string) ([]int, error) {
	input, err := ParseFile(filePath)

	if err != nil {
		return nil, err
	}

	numbers := make([]int, len(input))
	for i, line := range input {
		numbers[i], err = strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("error converting: %v to int, error: %w", line, err)
		}
	}

	return numbers, nil
}
