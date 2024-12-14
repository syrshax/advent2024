package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parse(filename string) []string {
	bytes, _ := os.ReadFile(filename)
	return strings.Split(string(bytes), "\n")
}

func main() {
	data := parse("text.txt")
	fmt.Println(data)
	var l, r []int
	for _, line := range data {
		nums := strings.Fields(line)
		if len(nums) == 2 {
			n1, _ := strconv.Atoi(nums[0])
			n2, _ := strconv.Atoi(nums[1])
			l = append(l, n1)
			r = append(r, n2)
		}
	}

	sort.Ints(l)
	sort.Ints(r)

	totalDistance := 0
	for i := range l {
		totalDistance += int(math.Abs(float64(l[i] - r[i])))
	}
	fmt.Println("Total distance:", totalDistance)
	score := similarityScore(l, r)
	fmt.Println("The total score is: ", score)
}

func similarityScore(a []int, b []int) int {
	countB := make(map[int]int)
	for _, v := range b {
		countB[v]++
	}

	score := 0
	for _, val := range a {
		score += val * countB[val]
	}
	return score
}
