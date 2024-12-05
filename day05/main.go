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

	incorrectUpdateMaps := []map[string]int{}
	incorrectUpdateSlices := [][]string{}

	sum := 0
	for index, updateMap := range updateMaps {
		updateSlice := updateSlices[index]
		valid := true

		for _, rule := range rules {
			if !checkRule(rule, updateMap) {
				incorrectUpdateMaps = append(incorrectUpdateMaps, updateMap)
				incorrectUpdateSlices = append(incorrectUpdateSlices, updateSlice)
				valid = false
				break
			}
		}

		if valid {
			sum += getMiddleElementAsInt(updateSlice)
		}
	}

	incorrectSum := 0
	for index, updateMap := range incorrectUpdateMaps {
		updateSlice := incorrectUpdateSlices[index]

		for {
			valid := true
			for _, rule := range rules {
				if !checkRule(rule, updateMap) {
					updateSlice = putBefore(updateSlice, updateMap[rule[1]]-1, updateMap[rule[0]]-1)
					updateMap = createIndexMap(updateSlice)
					valid = false
					break
				}
			}

			if valid {
				break
			}
		}

		incorrectSum += getMiddleElementAsInt(updateSlice)
	}

	fmt.Println("sum:         ", sum)
	fmt.Println("incorrectSum:", incorrectSum)
}

func setup() ([][]string, []map[string]int, [][]string) {
	input, _ := os.ReadFile("input.txt")
	inputs := strings.Split(string(input), "\n\n")

	ruleInput, updateInput := inputs[0], inputs[1]

	ruleLines := strings.Split(ruleInput, "\n")
	updateLines := strings.Split(updateInput, "\n")

	rules := [][]string{}
	for _, ruleLine := range ruleLines {
		rules = append(rules, strings.Split(ruleLine, "|"))
	}

	updateMaps := []map[string]int{}
	updateSlices := [][]string{}
	for _, updateLine := range updateLines {
		updateSlice := strings.Split(updateLine, ",")

		updateMaps = append(updateMaps, createIndexMap(updateSlice))
		updateSlices = append(updateSlices, updateSlice)
	}

	return rules, updateMaps, updateSlices
}

func createIndexMap(slice []string) map[string]int {
	indexMap := map[string]int{}

	for index, page := range slice {
		indexMap[page] = index + 1
	}

	return indexMap
}

func checkRule(rule []string, updateMap map[string]int) bool {
	return updateMap[rule[0]] == 0 ||
		updateMap[rule[1]] == 0 ||
		updateMap[rule[0]] < updateMap[rule[1]]
}

func putBefore(slice []string, x, y int) []string {
	valueY := slice[y]
	sliceWithoutY := append(slice[:y], slice[y+1:]...)
	return slices.Insert(sliceWithoutY, x, valueY)
}

func getMiddleElementAsInt(slice []string) int {
	result, _ := strconv.Atoi(slice[(len(slice)-1)/2])
	return result
}
