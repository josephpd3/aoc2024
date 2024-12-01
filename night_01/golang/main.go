package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Input struct {
	LeftList    []int
	RightList   []int
	RightCounts map[int]int
}

func getInput() (*Input, error) {
	input := &Input{
		RightCounts: make(map[int]int),
	}

	// Open input file
	inputFile, err := os.Open("../data/input.txt")
	if err != nil {
		return nil, err
	}

	// Read input file
	inputScanner := bufio.NewScanner(inputFile)
	for inputScanner.Scan() {
		inputLine := inputScanner.Text()

		// Skip empty lines
		if inputLine == "" {
			continue
		}

		// Parse input line
		inputNums := strings.Fields(inputLine)
		if len(inputNums) != 2 {
			return nil, errors.New("invalid input")
		}

		leftNum, err := strconv.Atoi(inputNums[0])
		if err != nil {
			return nil, err
		}
		input.LeftList = append(input.LeftList, leftNum)

		rightNum, err := strconv.Atoi(inputNums[1])
		if err != nil {
			return nil, err
		}
		input.RightList = append(input.RightList, rightNum)
		input.RightCounts[rightNum]++

	}

	sort.Ints(input.LeftList)
	sort.Ints(input.RightList)

	return input, nil
}

func absDifference(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func getDistance(input *Input) int {
	var total int
	for i := 0; i < len(input.LeftList); i++ {
		total += absDifference(input.LeftList[i], input.RightList[i])
	}
	return total
}

func getSimilarity(input *Input) int {
	var total int
	for i := 0; i < len(input.LeftList); i++ {
		leftNum := input.LeftList[i]
		total += input.RightCounts[leftNum] * leftNum
	}
	return total
}

func main() {
	input, err := getInput()
	if err != nil {
		log.Fatal(err)
	}
	distance := getDistance(input)
	fmt.Printf("Distance: %d\n", distance)
	similarity := getSimilarity(input)
	fmt.Printf("Similarity: %d\n", similarity)
}
