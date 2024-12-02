package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Reports [][]int
}

func getInput() (*Input, error) {
	input := &Input{}

	// Open input file
	inputFile, err := os.Open("../data/input.txt")
	if err != nil {
		return nil, err
	}

	inputScanner := bufio.NewScanner(inputFile)
	for inputScanner.Scan() {
		inputLine := inputScanner.Text()

		// Skip empty lines
		if inputLine == "" {
			continue
		}

		newReport := []int{}
		for _, num := range strings.Fields(inputLine) {
			n, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}
			newReport = append(newReport, n)
		}
		input.Reports = append(input.Reports, newReport)
	}

	return input, nil
}

func greaterThan(a, b int) bool {
	return a > b
}

func lessThan(a, b int) bool {
	return a < b
}

func absDifference(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func isSafe(report []int) bool {
	if len(report) == 1 {
		return true
	}

	firstNum := report[0]
	secondNum := report[1]

	var comparator func(int, int) bool

	if firstNum > secondNum {
		comparator = lessThan
	} else if firstNum < secondNum {
		comparator = greaterThan
	} else {
		return false
	}

	if absDifference(firstNum, secondNum) > 3 {
		return false
	}

	lastNum := secondNum

	for _, num := range report[2:] {
		if !comparator(num, lastNum) || absDifference(lastNum, num) > 3 || lastNum == num {
			return false
		}
		lastNum = num
	}

	return true
}

// isSafeWithDroppingOne is brute force bullshit but it works
// would not scale at ALL though if these reports were larger
func isSafeWithDroppingOne(report []int) bool {
	for i := 0; i < len(report); i++ {
		firstHalf := make([]int, len(report)-(len(report)-i))
		secondHalf := make([]int, len(report)-i-1)

		// Have to do this to avoid changing backing array
		copy(firstHalf, report[:i])
		copy(secondHalf, report[i+1:])

		if isSafe(append(firstHalf, secondHalf...)) {
			return true
		}
	}
	return false
}

func main() {
	input, err := getInput()
	if err != nil {
		panic(err)
	}

	totalSafe := 0

	for _, report := range input.Reports {
		if isSafe(report) {
			totalSafe++
		}
	}

	fmt.Println("Total safe:", totalSafe)

	totalSafeWithDroppingOne := 0

	for _, report := range input.Reports {
		if isSafeWithDroppingOne(report) {
			totalSafeWithDroppingOne++
		}
	}

	fmt.Println("Total safe with dropping one:", totalSafeWithDroppingOne)

	// first problem Test cases
	// fmt.Println(isSafe([]int{1, 2, 3, 4, 5}))    // true
	// fmt.Println(isSafe([]int{7, 6, 4, 2, 1}))    // true
	// fmt.Println(isSafe([]int{1, 2, 3, 4, 4}))    // false
	// fmt.Println(isSafe([]int{8, 6, 4, 4, 1}))    // false
	// fmt.Println(isSafe([]int{1, 2, 7, 8, 9}))    // false
	// fmt.Println(isSafe([]int{10, 10, 11, 8, 8})) // false

	// second problem Test cases
	// fmt.Println(isSafeWithDroppingOne([]int{7, 6, 4, 2, 1})) // true
	// fmt.Println(isSafeWithDroppingOne([]int{1, 2, 7, 8, 9})) // false
	// fmt.Println(isSafeWithDroppingOne([]int{9, 7, 6, 2, 1})) // false
	// fmt.Println(isSafeWithDroppingOne([]int{1, 3, 2, 4, 5})) // true
	// fmt.Println(isSafeWithDroppingOne([]int{8, 6, 4, 4, 1})) // true
	// fmt.Println(isSafeWithDroppingOne([]int{1, 3, 6, 7, 9})) // true
}
