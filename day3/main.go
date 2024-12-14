package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func enableMul(mulPos int, doIndices [][]int, dontIndices [][]int) bool {
	//init case, we got a dont, but mul is before, we enable the mul then
	if mulPos < dontIndices[0][0] {
		return true
	}
	lastDo := -1
	for i := len(doIndices) - 1; i >= 0; i-- {
		if doIndices[i][0] < mulPos {
			lastDo = doIndices[i][0]
			break
		}
	}
	lastDont := -1
	for i := len(dontIndices) - 1; i >= 0; i-- {
		if dontIndices[i][0] < mulPos {
			lastDont = dontIndices[i][0]
			break
		}
	}
	return lastDo > lastDont && mulPos > lastDo
}

func part1(text string) int {
	regexmul := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := regexmul.FindAllStringSubmatch(text, -1)
	sum := 0
	for _, match := range matches {
		n1, _ := strconv.Atoi(match[1])
		n2, _ := strconv.Atoi(match[2])
		sum += n1 * n2
	}
	return sum
}

func part2(text string) int {
	regexmul := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	do := regexp.MustCompile(`do\(\)`)
	dont := regexp.MustCompile(`don't\(\)`)

	matches := regexmul.FindAllStringSubmatch(text, -1)
	matchIndices := regexmul.FindAllStringSubmatchIndex(text, -1)
	doIndices := do.FindAllStringIndex(text, -1)
	dontIndices := dont.FindAllStringIndex(text, -1)

	sum := 0
	for i := 0; i < len(matches); i++ {
		if !enableMul(matchIndices[i][0], doIndices, dontIndices) {
			continue
		}
		n1, _ := strconv.Atoi(matches[i][1])
		n2, _ := strconv.Atoi(matches[i][2])
		sum += n1 * n2
	}
	return sum
}

func main() {
	data, _ := os.ReadFile(os.Args[1])
	content := strings.ReplaceAll(string(data), "\n", "")

	tot1 := part1(content)
	tot2 := part2(content)

	fmt.Printf("Part 1 tot: %d\n", tot1)
	fmt.Printf("Part 2 tot: %d\n", tot2)
}
