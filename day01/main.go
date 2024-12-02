package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	leftList := []int{}
	rightList := []int{}
	for _, line := range lines {
		numbers := strings.Split(line, "   ")

		leftNumber, _ := strconv.Atoi(numbers[0])
		leftList = appendSorted(leftList, leftNumber)

		rightNumber, _ := strconv.Atoi(numbers[1])
		rightList = appendSorted(rightList, rightNumber)
	}

	difference := 0
	rightMap := map[int]int{}
	for i := 0; i < len(leftList); i++ {
		difference = difference + abs(leftList[i]-rightList[i])
		rightMap[rightList[i]]++
	}

	similarity := 0
	for i := 0; i < len(leftList); i++ {
		similarity = similarity + (leftList[i] * rightMap[leftList[i]])
	}

	fmt.Printf("difference: %d\n", difference)
	fmt.Printf("similarity: %d\n", similarity)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func appendSorted(inputList []int, num int) []int {
	listLength := len(inputList)

	if listLength == 0 || inputList[listLength-1] <= num {
		return append(inputList, num)
	}

	if inputList[0] >= num {
		return append([]int{num}, inputList...)
	}

	for i := 0; i < listLength; i++ {
		if inputList[i] >= num {
			return slices.Insert(inputList, i, num)
		}
	}

	return append(inputList, num)
}
