package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Equation struct {
	target  int
	numbers []int
}

func parse(line string) (Equation, error) {
	parts := strings.Split(line, ": ")
	if len(parts) != 2 {
		return Equation{}, fmt.Errorf("invalid format sir: %s", line)
	}

	target, err := strconv.Atoi(parts[0])
	if err != nil {
		return Equation{}, err
	}

	numStrs := strings.Fields(parts[1])
	numbers := make([]int, len(numStrs))
	for i, ns := range numStrs {
		numbers[i], err = strconv.Atoi(ns)
		if err != nil {
			return Equation{}, err
		}
	}

	return Equation{target: target, numbers: numbers}, nil
}

func evaluateExpression(numbers []int, operators []string) int {
	result := numbers[0]
	for i := 0; i < len(operators); i++ {
		switch operators[i] {
		case "+":
			result += numbers[i+1]
		case "*":
			result *= numbers[i+1]
		}
	}
	return result
}

func generateOperatorCombinations(pos int, operators []string, results *[][]string) {
	if pos == len(operators) {
		opsCopy := make([]string, len(operators))
		copy(opsCopy, operators)
		*results = append(*results, opsCopy)
		return
	}

	for _, op := range []string{"+", "*"} {
		operators[pos] = op
		generateOperatorCombinations(pos+1, operators, results)
	}
}

func canSolveEquation(eq Equation) bool {
	if len(eq.numbers) == 1 {
		return eq.target == eq.numbers[0]
	}

	operators := make([]string, len(eq.numbers)-1)
	var possibleOps [][]string
	generateOperatorCombinations(0, operators, &possibleOps)

	for _, ops := range possibleOps {
		if evaluateExpression(eq.numbers, ops) == eq.target {
			return true
		}
	}
	return false
}

func main() {
	t := time.Now()
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	sum := 0

	for _, line := range lines {
		eq, err := parse(line)
		if err != nil {
			fmt.Printf("Error parsing equation '%s': %v\n", line, err)
			continue
		}

		if canSolveEquation(eq) {
			sum += eq.target
		}
	}
	tottime := time.Now().Sub(t)
	fmt.Println("time for benchmark", tottime)
	fmt.Printf("tot calibration: %d\n", sum)

}
