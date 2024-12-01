package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"strconv"
	"strings"
)

func getScanner() (*bufio.Scanner, *os.File, error) {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %w", err)
	}

	// Create a scanner
	scanner := bufio.NewScanner(file)
	return scanner, file, nil
}

func parseFile() ([]int, []int, map[int]int, map[int]int, error) {
	// Initialize slices and maps
	var c1Slice []int
	var c2Slice []int
	c1Counts := make(map[int]int)
	c2Counts := make(map[int]int)

	// Get the scanner
	scanner, file, err := getScanner()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	defer file.Close()

	// Read and parse the file
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		// Ensure at least two fields are present
		if len(fields) >= 2 {
			// Convert to integers
			num1, err1 := strconv.Atoi(fields[0])
			num2, err2 := strconv.Atoi(fields[1])
			if err1 != nil || err2 != nil {
				fmt.Println("Error converting to integer:", err1, err2)
				continue
			}

			// Populate slices
			c1Slice = append(c1Slice, num1)
			c2Slice = append(c2Slice, num2)

			// Populate maps
			c1Counts[num1]++
			c2Counts[num2]++
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, nil, nil, nil, err
	}
    sort.Ints(c1Slice)
    sort.Ints(c2Slice)

	return c1Slice, c2Slice, c1Counts, c2Counts, nil
}

func partOne(s []int, l []int) int {
    sum := 0
    for i := range s {
        diff := l[i] - s[i]
        if diff < 0 {
            diff = -diff
        }
        sum += diff
    }
    return sum
}

func partTwo(s []int, m map[int]int) int {
    sum := 0
    for i := range s {
        // fmt.Printf("The num at this position is %v and its count is %v", s[i], m[i])
        sum += s[i] * m[s[i]]
    }
    return sum
}

func main() {

    c1Slice, c2Slice, _, c2Counts, err := parseFile()
    if err != nil {
        fmt.Printf("Error %v", err)
        return
    }

    fmt.Printf("Part 1 Answer: %v\n", partOne(c1Slice, c2Slice))
    fmt.Printf("Part 2 Answer: %v\n", partTwo(c1Slice, c2Counts))

}
