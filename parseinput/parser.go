package parseinput

import (
	"bufio"
	"os"
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
