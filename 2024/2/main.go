package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file", err)
	}

	scanner := bufio.NewScanner(file)

	inputVals := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		subSlice := []int{}
		for _, v := range strings.Split(line, " ") {
			curVal, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal("Error converting string to integer", err)
			}
			subSlice = append(subSlice, curVal)
		}
		inputVals = append(inputVals, subSlice)
		// fmt.Printf("This is a line from the input file: %s\n", line)
	}
	fmt.Printf("Part one: %d", partOne(inputVals))
	fmt.Printf("Part two: %d", partTwo(inputVals))

}

func partOne(s [][]int) int {
	validCount := 0

	for _, line := range s {
		if len(line) < 2 {
			continue // Skip lines too short to be valid
		}

		isIncreasing := false
		isDecreasing := false
		valid := true

		for i := 1; i < len(line); i++ {
			diff := line[i] - line[i-1]

			// Check if the difference is invalid
			if diff < -3 || diff > 3 || diff == 0 {
				valid = false
				break
			}

			// Determine if it's increasing or decreasing
			if diff > 0 {
				isIncreasing = true
			} else if diff < 0 {
				isDecreasing = true
			}

			// If both increasing and decreasing are detected, it's invalid
			if isIncreasing && isDecreasing {
				valid = false
				break
			}
		}

		if valid {
			validCount++
		}
	}

	return validCount
}

func partTwo (s [][]int) int {
	validCount := 0

	isSafe := func(line []int) bool {
		if len(line) < 2 {
			return false // Too short to be valid
		}

		isIncreasing := false
		isDecreasing := false

		for i := 1; i < len(line); i++ {
			diff := line[i] - line[i-1]

			// Check if the difference is invalid
			if diff < -3 || diff > 3 || diff == 0 {
				return false
			}

			// Determine if it's increasing or decreasing
			if diff > 0 {
				isIncreasing = true
			} else if diff < 0 {
				isDecreasing = true
			}

			// If both increasing and decreasing are detected, it's invalid
			if isIncreasing && isDecreasing {
				return false
			}
		}
		return true
	}

	for _, line := range s {
		// Check if the report is already safe
		if isSafe(line) {
			validCount++
			continue
		}

		// Check if removing a single level makes the report safe
		for i := 0; i < len(line); i++ {
			// Create a new line with the i-th level removed
			modifiedLine := append([]int{}, line[:i]...)
			modifiedLine = append(modifiedLine, line[i+1:]...)

			if isSafe(modifiedLine) {
				validCount++
				break
			}
		}
	}

	return validCount
}
