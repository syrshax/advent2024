package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Condition struct {
	before, after int
}

func main() {
	input, err := os.ReadFile("test_input.txt")
	if err != nil {
		fmt.Printf("error parsing file xxx\n")
		os.Exit(1)
		return
	}

	//viendo el file damos cuenta que esta separado por doble \n\n
	textpart := strings.Split(string(input), "\n\n")
	var condition []Condition
	sum := 0

	for _, line := range strings.Split(textpart[0], "\n") {
		parts := strings.Split(line, "|")
		before, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])
		condition = append(condition, Condition{before, after})
	}

	for _, line := range strings.Split(textpart[1], "\n") {
		var sec []int
		for _, numStr := range strings.Split(line, ",") {
			num, _ := strconv.Atoi(numStr)
			sec = append(sec, num)
		}
		positions := make(map[int]int)
		for i, num := range sec {
			positions[num] = i
		}
		valid := true
		for _, con := range condition {
			before, beforeExists := positions[con.before]
			after, afterExists := positions[con.after]
			if beforeExists && afterExists && before > after {
				valid = false
				break
			}
		}
		if valid {
			sum += sec[len(sec)/2]
		}
	}
	fmt.Printf("sum: %d", sum)
}
