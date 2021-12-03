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
}

func partOne(input []string) {
	bitArray, err := createBitArray(input)

	if err != nil {
		panic(err)
	}

	reverseBitArray := make([]byte, len(bitArray))
	for i, bit := range bitArray {
		reverseBitArray[i] = reverseBit(bit)
	}
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

func createBitArray(input []string) ([]byte, error) {
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
		go countBits(byteChan, returnChan)

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

func reverseBit(x byte) byte {
	if x == '1' {
		return '0'
	} else {
		return '1'
	}
}

func countBits(ch chan byte, returnChan chan byte) {
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
	} else {
		returnChan <- '0'
	}
}
