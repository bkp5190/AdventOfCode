package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"

	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file", err)
	}

	scanner := bufio.NewScanner(file)

	inputVals := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		inputVals = append(inputVals, line)
	}

	fmt.Printf("Part one: %d\n", partOne(inputVals))
	fmt.Printf("Part two: %d\n", partTwo(inputVals))
}

func partOne(s []string) int {
	// Apply a regex to capture all occurrences of mul(int, int)
	// "mul\([0-9][0-9][0-9],[0-9][0-9][0-9]\)"
	re, _ := regexp.Compile("mul\\([0-9][0-9]?[0-9]?,[0-9][0-9]?[0-9]?\\)")
	sum := 0
	for i := range s {

		matches := re.FindAllString(s[i], -1)

		for _, v := range matches {
			v = strings.Replace(v, "mul(", "", -1)
			v = strings.Replace(v, ")", "", -1)
			vals := strings.Split(v, ",")

			// Ensure we have exactly two values to process (pairs of values)
			if len(vals) == 2 {
				prevVal, err := strconv.Atoi(vals[0])
				if err != nil {
					fmt.Println("Error converting prevVal:", err)
					continue
				}
				curVal, err := strconv.Atoi(vals[1])
				if err != nil {
					fmt.Println("Error converting curVal:", err)
					continue
				}
				sum += prevVal * curVal
			} else {
				fmt.Println("Skipping invalid pair:", vals)
			}
		}
	}
	return sum
}

func partTwo(s []string) int {
    // Compile regular expressions for do(), don't(), and mul()
    doRe, _ := regexp.Compile("do\\(\\)")
    dontRe, _ := regexp.Compile("don't\\(\\)")
    mulRe, _ := regexp.Compile("mul\\([0-9]+,[0-9]+\\)")

    sum := 0
    mulEnabled := true // At the start, mul() instructions are enabled

    for _, line := range s {
        // Find all matches of each regex
        doIndices := doRe.FindAllStringIndex(line, -1)
        dontIndices := dontRe.FindAllStringIndex(line, -1)
        mulIndices := mulRe.FindAllStringIndex(line, -1)
        mulMatches := mulRe.FindAllString(line, -1)

        // Collect all events in chronological order
        events := []struct {
            index int
            kind  string
            value string
        }{}

        for _, idx := range doIndices {
            events = append(events, struct {
                index int
                kind  string
                value string
            }{index: idx[0], kind: "do", value: ""})
        }

        for _, idx := range dontIndices {
            events = append(events, struct {
                index int
                kind  string
                value string
            }{index: idx[0], kind: "don't", value: ""})
        }

        for i, idx := range mulIndices {
            events = append(events, struct {
                index int
                kind  string
                value string
            }{index: idx[0], kind: "mul", value: mulMatches[i]})
        }

        // Sort events by their index
        sort.Slice(events, func(i, j int) bool {
            return events[i].index < events[j].index
        })

        // Process events in order
        for _, event := range events {
            switch event.kind {
            case "do":
                mulEnabled = true
            case "don't":
                mulEnabled = false
            case "mul":
                if mulEnabled {
                    v := event.value
                    v = strings.Replace(v, "mul(", "", -1)
                    v = strings.Replace(v, ")", "", -1)
                    vals := strings.Split(v, ",")
                    if len(vals) == 2 {
                        a, err1 := strconv.Atoi(vals[0])
                        b, err2 := strconv.Atoi(vals[1])
                        if err1 == nil && err2 == nil {
                            sum += a * b
                        }
                    }
                }
            }
        }
    }

    return sum
}
