package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	guard    = "^"
	wall     = "#"
	visited  = "X"
	obstacle = "O"
)

func main() {
	guardMap, gx, gy := setup()

	sum := 0
	angle := math.Pi
	dx, dy := calcDirection(angle)
	for {
		if guardMap[gx][gy] != visited {
			sum++
			guardMap[gx][gy] = visited
		}

		if !isValid(guardMap, gx+dx, gy+dy) {
			break
		}

		if guardMap[gx+dx][gy+dy] == wall {
			angle = nextAngle(angle)
			dx, dy = calcDirection(angle)
		}

		gx += dx
		gy += dy
	}

	fmt.Println("sum:", sum)
}

func setup() ([][]string, int, int) {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	guardMap := [][]string{}
	var gx, gy int

	for index, line := range lines {
		guardMap = append(guardMap, strings.Split(line, ""))
		if strings.Contains(line, guard) {
			gx = index
			gy = strings.Index(line, guard)
		}
	}

	return guardMap, gx, gy
}

func isValid(array [][]string, x, y int) bool {
	return x >= 0 && x < len(array) && y >= 0 && y < len(array[x])
}

func nextAngle(angle float64) float64 {
	return math.Mod(angle+1.5*math.Pi, 2*math.Pi)
}

func calcDirection(angle float64) (int, int) {
	return int(math.Cos(angle)), int(math.Sin(angle))
}
