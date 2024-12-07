package main

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"slices"
	"strings"
)

const (
	guard       = "^"
	wall        = "#"
	visited     = "X"
	visitedLoop = "Y"
	obstacle    = "O"
)

func main() {
	guardMap, gx, gy := setup("test.txt")

	sum, loops := dfs(guardMap, gx, gy, math.Pi)

	fmt.Println("sum:  ", sum)
	fmt.Println("loops:", len(loops))
}

func setup(file string) ([][]string, int, int) {
	input, _ := os.ReadFile(file)
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

func dfs(array [][]string, x, y int, angle float64) (int, []map[string]struct{}) {
	count := 0
	loops := []map[string]struct{}{}
	dx, dy := calcDirection(angle)

	if array[x][y] != visited {
		array[x][y] = visited
		count = 1
	}

	if !isValid(array, x+dx, y+dy) {
		return count, loops
	}

	if array[x+dx][y+dy] == wall {
		angle = nextAngle(angle)
		dx, dy = calcDirection(angle)
	}

	arrayCopy := copy(array)
	arrayCopy[x+dx][y+dy] = obstacle
	parents := createEmptyArrayFrom(arrayCopy)
	loop := dfsLoop(arrayCopy, parents, x, y, angle)

	prevCount, prevLoops := dfs(array, x+dx, y+dy, angle)

	loops = prevLoops
	if len(loop) != 0 && !arrayContains(loops, loop) {
		loops = append(loops, loop)
	}

	return prevCount + count, loops
}

func dfsLoop(array [][]string, parent [][][]int, x, y int, angle float64) map[string]struct{} {
	dx, dy := calcDirection(angle)

	if array[x][y] != visitedLoop {
		array[x][y] = visitedLoop
	}

	if !isValid(array, x+dx, y+dy) {
		return map[string]struct{}{}
	}

	if array[x+dx][y+dy] == wall || array[x+dx][y+dy] == obstacle {
		angle = nextAngle(angle)
		dx, dy = calcDirection(angle)
	}

	if array[x+dx][y+dy] == visitedLoop && slices.Equal(parent[x+dx][y+dy], []int{x, y}) {
		result := map[string]struct{}{}
		for i, row := range parent {
			for j, col := range row {
				if len(col) > 0 {
					result[fmt.Sprintf("[%d,%d,%d,%d]", i, j, col[0], col[1])] = struct{}{}
				}
			}
		}
		return result
	}

	parent[x+dx][y+dy] = []int{x, y}
	return dfsLoop(array, parent, x+dx, y+dy, angle)
}

func printArray(array [][]string) {
	for _, line := range array {
		for _, elem := range line {
			fmt.Print(elem)
		}
		fmt.Println()
	}
	fmt.Println()
}

func copy(array [][]string) [][]string {
	copy := [][]string{}
	for _, elem := range array {
		copy = append(copy, append([]string{}, elem...))
	}

	return copy
}

func createEmptyArrayFrom(array [][]string) [][][]int {
	emptyArray := [][][]int{}
	for i := 0; i < len(array); i++ {
		emptyArray = append(emptyArray, [][]int{})
		for j := 0; j < len(array[i]); j++ {
			emptyArray[i] = append(emptyArray[i], []int{})
		}
	}

	return emptyArray
}

func arrayContains(array []map[string]struct{}, new map[string]struct{}) bool {
	for i := 0; i < len(array); i++ {
		if reflect.DeepEqual(array[i], new) {
			return true
		}
	}

	return false
}
