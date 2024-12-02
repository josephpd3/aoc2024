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

func isSafelyMonotonic(report []int, freebies int, direction int) bool {
	if len(report) == 1 {
		return true
	}

	freebiesUsed := 0

	var comparator func(int, int) bool

	if direction > 0 {
		comparator = greaterThan
	} else if direction < 0 {
		comparator = lessThan
	} else {
		panic("direction must be non-zero")
	}

	lastNum := report[0]

	for _, num := range report[1:] {
		if !comparator(num, lastNum) || absDifference(lastNum, num) > 3 || lastNum == num {
			if freebiesUsed < freebies {
				freebiesUsed++
				continue
			} else {
				return false
			}
		}
		lastNum = num
	}

	return true
}

// Okay here's where it gets weird and I am making assumptions
//
// The report is safe if either:
// 1. The report minus the first entry is safe with no modifications
// 2. The report is safe minus one other entry or minus no entries
//
// From there you can just check if the report is safely monotonic in either direction
func isSafeWithDroppingOne(report []int) bool {
	reportWithoutFirst := make([]int, len(report)-1)
	copy(reportWithoutFirst, report[1:])

	safeWithoutFirst := (
		isSafelyMonotonic(reportWithoutFirst, 0, 1) || isSafelyMonotonic(reportWithoutFirst, 0, -1))

	return safeWithoutFirst || isSafelyMonotonic(report, 1, 1) || isSafelyMonotonic(report, 1, -1)
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

	fmt.Println("Total safe:", totalSafe) // should be 390

	totalSafeWithDroppingOne := 0

	for _, report := range input.Reports {
		if isSafeWithDroppingOne(report) {
			totalSafeWithDroppingOne++
		}
	}

	// should be 439
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
