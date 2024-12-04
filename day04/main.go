package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	word  = "XMAS"
	xWord = "MAS"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")
	letters := [][]string{}
	for _, line := range lines {
		letters = append(letters, strings.Split(line, ""))
	}

	wordCount := 0
	xWordCount := 0
	for i := 0; i < len(letters); i++ {
		for j := 0; j < len(letters[i]); j++ {
			if (letters[i][j]) == string(word[0]) {
				wordCount = wordCount + countWords(letters, i, j)
			}

			if (letters[i][j]) == string(xWord[1]) && hasXWord(letters, i, j) {
				xWordCount++
			}
		}
	}

	fmt.Println("wordCount: ", wordCount)
	fmt.Println("xWordCount:", xWordCount)
}

func countWords(letters [][]string, x, y int) int {
	count := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !isValid(letters, x+i, y+j) {
				continue
			}

			condition := true
			for k := 1; k < len(word) && condition; k++ {
				condition = condition &&
					isValid(letters, k*i+x, k*j+y) &&
					letters[k*i+x][k*j+y] == string(word[k])
			}

			if !condition {
				continue
			}

			count++
		}
	}

	return count
}

func hasXWord(letters [][]string, x, y int) bool {
	xMap := map[string]int{}
	xRange := []int{-1, 1}

	for _, k := range xRange {
		for _, l := range xRange {
			if isValid(letters, x+k, y+l) {
				xMap[letters[x+k][y+l]]++
			}
		}
	}

	condition := xMap[string(xWord[0])] == 2 &&
		xMap[string(xWord[2])] == 2 &&
		letters[x-1][y-1] != letters[x+1][y+1]

	return condition
}

func isValid(array [][]string, x, y int) bool {
	return x >= 0 && x < len(array) && y >= 0 && y < len(array[x])
}
