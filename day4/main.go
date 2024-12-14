package main

import (
	"fmt"
	"os"
	"strings"
)

func instantmapPositionsLookup(i, j int) [][3][2]int {
	return [][3][2]int{
		{{i - 1, j}, {i - 2, j}, {i - 3, j}},
		{{i + 1, j}, {i + 2, j}, {i + 3, j}},
		{{i, j + 1}, {i, j + 2}, {i, j + 3}},
		{{i, j - 1}, {i, j - 2}, {i, j - 3}},
		{{i - 1, j + 1}, {i - 2, j + 2}, {i - 3, j + 3}},
		{{i - 1, j - 1}, {i - 2, j - 2}, {i - 3, j - 3}},
		{{i + 1, j + 1}, {i + 2, j + 2}, {i + 3, j + 3}},
		{{i + 1, j - 1}, {i + 2, j - 2}, {i + 3, j - 3}},
	}
}

func findXMAS(grid []string) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] != 'X' {
				continue
			}
			mapPossiblePositions := instantmapPositionsLookup(i, j)
			for _, mapPositions := range mapPossiblePositions {
				word := "X"
				valid := true
				for _, pos := range mapPositions {
					row, col := pos[0], pos[1]
					if row < 0 || row >= rows || col < 0 || col >= cols {
						valid = false
						break
					}
					word += string(grid[row][col])
				}
				if valid && word == "XMAS" {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	lines := strings.Fields(string(data))
	count := findXMAS(lines)
	fmt.Println("TOTAL XMAS: ===> ", count)
}
