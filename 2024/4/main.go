package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type offset struct {
    x int
    y int
}

func main() {

    dirs := map[string]offset {
        "north": offset{x: 0,y: 1},
        "south": offset{x: 0, y: -1},
        "east": offset{x: 1, y: 0},
        "west": offset{x: -1, y: 0},
        "northeast": offset{x: 1, y: 1},
        "southeast": offset{x: 1, y: -1},
        "southwest": offset{x: -1, y: -1},
        "northwest": offset{x: -1, y: 1},
    }

    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal("Error opening file", err)
    }

    search := make([][]rune, 0)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text() // Use Text() to directly get a string
        tmpArr := make([]rune, 0)
        for _, ch := range line {
            tmpArr = append(tmpArr, ch) // Convert each character to a rune
        }
        search = append(search, tmpArr)
    }

	fmt.Printf("Part one: %d\n", partOne(search, dirs))
	fmt.Printf("Part two: %d\n", partTwo(search))

}

func partOne(search [][]rune, dirs map[string]offset) int {

    wordCount := 0
    word := "XMAS"

	for x, row := range search {
		for y, char := range row {
			if char == rune(word[0]) { // Start search if the first letter matches
				for _, dir := range dirs {
					found := true
					for i := 1; i < len(word); i++ {
						newX, newY := x+(i*dir.x), y+(i*dir.y)
						// Check if within bounds and matches the next character
						if newX < 0 || newY < 0 || newX >= len(search) || newY >= len(row) ||
							search[newX][newY] != rune(word[i]) {
							found = false
							break
						}
					}
					if found {
						wordCount++
					}
				}
			}
		}
	}
    return wordCount
}

func partTwo(m [][]rune) int {

    total := 0
    rows := len(m)
    if rows == 0 {
        return total
    }
    cols := len(m[0])

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if m[i][j] != 'A' {
                continue
            }
            // Check the conditions for the four patterns around 'A'
            if i-1 >= 0 && j-1 >= 0 && i+1 < rows && j+1 < cols {
                if m[i-1][j-1] == 'M' && m[i-1][j+1] == 'M' && m[i+1][j-1] == 'S' && m[i+1][j+1] == 'S' {
                    total++
                }
                if m[i-1][j-1] == 'M' && m[i-1][j+1] == 'S' && m[i+1][j-1] == 'M' && m[i+1][j+1] == 'S' {
                    total++
                }
                if m[i-1][j-1] == 'S' && m[i-1][j+1] == 'M' && m[i+1][j-1] == 'S' && m[i+1][j+1] == 'M' {
                    total++
                }
                if m[i-1][j-1] == 'S' && m[i-1][j+1] == 'S' && m[i+1][j-1] == 'M' && m[i+1][j+1] == 'M' {
                    total++
                }
            }
        }
    }

    return total
}
