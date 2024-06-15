package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	output := spiralOrder(input)
	fmt.Println(output)
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	var result []int
	x, y := 0, 0
	flow := "right"
	visitedMap := map[string]bool{}
	numberOfMatrixElements := len(matrix) * len(matrix[0])
	for i := 0; i < numberOfMatrixElements; i++ {
		result = append(result, matrix[x][y])
		keystring := strconv.Itoa(x) + "-" + strconv.Itoa(y)
		visitedMap[keystring] = true
		nextX, nextY := generateNextXY(x, y, flow)

		keystring = strconv.Itoa(nextX) + "-" + strconv.Itoa(nextY)
		if _, visited := visitedMap[keystring]; visited ||
			nextX >= len(matrix) ||
			nextY >= len(matrix[0]) ||
			nextX < 0 ||
			nextY < 0 {
			flow = changeFlow(flow)
			nextX, nextY = generateNextXY(x, y, flow)
			keystring = strconv.Itoa(nextX) + "-" + strconv.Itoa(nextY)
			if _, visited := visitedMap[keystring]; visited ||
				nextX >= len(matrix) ||
				nextY >= len(matrix[0]) ||
				nextX < 0 ||
				nextY < 0 {
				break
			}
		}

		x, y = nextX, nextY
	}

	return result
}

func generateNextXY(x, y int, flow string) (int, int) {
	nextX, nextY := x, y
	if flow == "right" {
		nextX, nextY = x, y+1
	} else if flow == "bottom" {
		nextX, nextY = x+1, y
	} else if flow == "left" {
		nextX, nextY = x, y-1
	} else if flow == "top" {
		nextX, nextY = x-1, y
	}
	return nextX, nextY
}

func changeFlow(flow string) string {
	switch flow {
	case "right":
		return "bottom"
	case "bottom":
		return "left"
	case "left":
		return "top"
	case "top":
		return "right"
	}

	return flow
}
