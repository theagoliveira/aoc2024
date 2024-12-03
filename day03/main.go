package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	strInput := strings.ReplaceAll(string(input), "\n", "")

	mulRegex, _ := regexp.Compile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	submatches := mulRegex.FindAllStringSubmatch(strInput, -1)

	condInput := strInput
	for _, condRegexStr := range []string{`don't\(\).*?do\(\)`, `don't\(\).*$`} {
		condRegex, _ := regexp.Compile(condRegexStr)
		condInput = condRegex.ReplaceAllLiteralString(condInput, "")
	}
	condSubmatches := mulRegex.FindAllStringSubmatch(condInput, -1)

	fmt.Println("sum:    ", submatchSum(submatches))
	fmt.Println("condSum:", submatchSum(condSubmatches))
}

func submatchSum(submatches [][]string) int {
	sum := 0

	for _, submatch := range submatches {
		num1, _ := strconv.Atoi(submatch[1])
		num2, _ := strconv.Atoi(submatch[2])
		sum = sum + (num1 * num2)
	}

	return sum
}
