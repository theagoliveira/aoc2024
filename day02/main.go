package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	reports := strings.Split(string(input), "\n")

	safe := 0
	safeWithDampener := 0
	for _, report := range reports {
		strLevels := strings.Split(report, " ")

		levels := []int{}
		for i := 0; i < len(strLevels); i++ {
			intLevel, _ := strconv.Atoi(strLevels[i])
			levels = append(levels, intLevel)
		}

		if isSafe(levels) {
			safe++
		}

		if isSafeWithDampener(levels) {
			safeWithDampener++
		}
	}

	fmt.Println("safe:            ", safe)
	fmt.Println("safeWithDampener:", safeWithDampener)
}

func isSafe(levels []int) bool {
	refDir := levels[0] > levels[1]

	for i := 0; i < len(levels)-1; i++ {
		curr := levels[i]
		next := levels[i+1]

		dir := curr > next
		if dir != refDir {
			return false
		}

		diff := abs(curr - next)
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func isSafeWithDampener(levels []int) bool {
	for i := 0; i < len(levels); i++ {
		variation := append(copyIntSlice(levels[:i]), copyIntSlice(levels[i+1:])...)
		if isSafe(variation) {
			return true
		}
	}

	return false
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func copyIntSlice(slice []int) []int {
	return append([]int{}, slice...)
}
