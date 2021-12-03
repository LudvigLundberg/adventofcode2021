package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/LudvigLundberg/adventofcode2021/parseinput"
)

func main() {
	input, err := parseinput.ParseFile("input")

	if err != nil {
		log.Fatalf("error parsing input: %v", err)
	}

	fmt.Println("--------- Part One --------")
	partOne(input)
	fmt.Println("--------- Part Two --------")
	partTwo(input)
}

func partOne(input []string) {
	bitArray, err := createBitArray(input, '1')

	if err != nil {
		panic(err)
	}

	reverseBitArray := reverseBitArray(bitArray)

	gamma, err := strconv.ParseInt(string(bitArray), 2, 64)
	if err != nil {
		panic(err)
	}

	epsilon, err := strconv.ParseInt(string(reverseBitArray), 2, 64)

	if err != nil {
		panic(err)
	}

	fmt.Printf("gamma: %d, epsilon: %d, multiplied: %d \n", gamma, epsilon, gamma*epsilon)
}

func partTwo(input []string) {
	bitArray, err := createBitArray(input, '1')
	if err != nil {
		panic(err)
	}

	reverseBitArray := reverseBitArray(bitArray)

	oxygenString, err := filterCriteria(input, bitArray, 0, bitArray[0], '1')
	if err != nil {
		panic(err)
	}

	co2String, err := filterCriteria(input, reverseBitArray, 0, reverseBitArray[0], '0')
	if err != nil {
		panic(err)
	}

	oxygen, err := strconv.ParseInt(oxygenString, 2, 64)

	if err != nil {
		panic(err)
	}

	co2, err := strconv.ParseInt(co2String, 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Printf("oxygen: %d, co2: %d, multiplied: %d\n", oxygen, co2, oxygen*co2)
}

func filterCriteria(input []string, criteria []byte, index int, bit byte, priority byte) (string, error) {
	if index >= len(criteria) {
		return "", fmt.Errorf("did not find any matching strings")
	}

	filtered := filterByBit(input, bit, index)
	if len(filtered) == 1 {
		return filtered[0], nil
	} else {
		criteria, err := createBitArray(filtered, '1')
		if err != nil {
			return "", err
		}
		if priority == '0' {
			criteria = reverseBitArray(criteria)
		}
		return filterCriteria(filtered, criteria, index+1, criteria[index+1], priority)
	}
}

func reverseBitArray(bits []byte) []byte {
	reverse := make([]byte, len(bits))
	for i, bit := range bits {
		reverse[i] = reverseBit(bit)
	}
	return reverse
}

func createBitArray(input []string, priority byte) ([]byte, error) {
	var nrOfBits int
	if len(input) > 0 {
		nrOfBits = len(input[0])
	}

	bitArray := make([]byte, nrOfBits)

	var wg sync.WaitGroup
	wg.Add(nrOfBits)

	for i := range bitArray {
		byteChan := make(chan byte, len(input))
		returnChan := make(chan byte, 1)
		go countBits(byteChan, returnChan, priority)

		go func(i int) {
			for j := range input {
				byteChan <- input[j][i]
			}
			close(byteChan)
			bitArray[i] = <-returnChan
			wg.Done()
		}(i)
	}

	wg.Wait()

	return bitArray, nil
}

func filterByBit(input []string, bit_criteria byte, index int) []string {
	output := make([]string, 0, len(input))

	for _, line := range input {
		if line[index] == bit_criteria {
			output = append(output, line)
		}
	}

	return output
}

func reverseBit(x byte) byte {
	if x == '1' {
		return '0'
	} else {
		return '1'
	}
}

func countBits(ch chan byte, returnChan chan byte, priority byte) {
	count := 0
	defer close(returnChan)
	for input := range ch {
		switch input {
		case '1':
			count++
		case '0':
			count--
		}
	}

	if count > 0 {
		returnChan <- '1'
	} else if count < 0 {
		returnChan <- '0'
	} else {
		returnChan <- priority
	}
}
