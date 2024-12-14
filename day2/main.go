package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parse(filename string) []string {
	bytes, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	return lines
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func toNumbers(s []string) [][]int {
	result := make([][]int, len(s))
	for i, line := range s {
		numStrings := strings.Fields(line)
		nums := make([]int, len(numStrings))
		for j, numStr := range numStrings {
			nums[j], _ = strconv.Atoi(numStr)
		}
		result[i] = nums
	}
	return result
}

func safeSlice(row []int) bool {
	increasing, decreasing := true, true
	for i := 0; i < len(row)-1; i++ {
		diff := row[i+1] - row[i]
		increasing = increasing && diff > 0
		decreasing = decreasing && diff < 0
		if abs(diff) < 1 || abs(diff) > 3 || (!increasing && !decreasing) {
			return false
		}
	}
	return true
}

func resolve(numbers [][]int, tolerate bool) int {
	safeLevelCounter := 0
	for _, row := range numbers {
		if safeSlice(row) {
			safeLevelCounter++
			continue
		}

		if tolerate {
			for i := range row {
				newRow := append(append([]int{}, row[:i]...), row[i+1:]...)
				if safeSlice(newRow) {
					safeLevelCounter++
					break
				}
			}
		}
	}
	return safeLevelCounter
}

func main() {
	data := parse("input.txt")
	numbers := toNumbers(data)

	fmt.Printf("Part 1: %d\n", resolve(numbers, false))
	fmt.Printf("Part 2: %d\n", resolve(numbers, true))
}
