package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	rules, updateMaps, updateSlices := setup()

	sum := 0
	incorrectSum := 0
	for index, updateMap := range updateMaps {
		updateSlice := updateSlices[index]
		valid := true

		for _, rule := range rules {
			if !checkRule(rule, updateMap) {
				slices.SortFunc(updateSlice, func(a, b int) int {
					return compareRules(rules, a, b)
				})
				incorrectSum += getMiddleElement(updateSlice)
				valid = false
				break
			}
		}

		if valid {
			sum += getMiddleElement(updateSlice)
		}
	}

	fmt.Println("sum:         ", sum)
	fmt.Println("incorrectSum:", incorrectSum)
}

func setup() ([][]int, []map[int]int, [][]int) {
	input, _ := os.ReadFile("input.txt")
	inputs := strings.Split(string(input), "\n\n")

	ruleInput, updateInput := inputs[0], inputs[1]

	ruleLines := strings.Split(ruleInput, "\n")
	updateLines := strings.Split(updateInput, "\n")

	rules := [][]int{}
	for _, ruleLine := range ruleLines {
		rules = append(rules, convertToInt(strings.Split(ruleLine, "|")))
	}

	updateMaps := []map[int]int{}
	updateSlices := [][]int{}
	for _, updateLine := range updateLines {
		updateSlice := convertToInt(strings.Split(updateLine, ","))

		updateMaps = append(updateMaps, createIndexMap(updateSlice))
		updateSlices = append(updateSlices, updateSlice)
	}

	return rules, updateMaps, updateSlices
}

func convertToInt(array []string) []int {
	nums := []int{}
	for _, str := range array {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}
	return nums
}

func createIndexMap(slice []int) map[int]int {
	indexMap := map[int]int{}

	for index, page := range slice {
		indexMap[page] = index + 1
	}

	return indexMap
}

func checkRule(rule []int, updateMap map[int]int) bool {
	return updateMap[rule[0]] == 0 ||
		updateMap[rule[1]] == 0 ||
		updateMap[rule[0]] < updateMap[rule[1]]
}

func getMiddleElement(slice []int) int {
	return slice[(len(slice)-1)/2]
}

func compareRules(rules [][]int, a, b int) int {
	for _, rule := range rules {
		if rule[0] == a && rule[1] == b {
			return -1
		} else if rule[0] == b && rule[1] == a {
			return 1
		}
	}

	return compare(a, b)
}

func compare(a, b int) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}
