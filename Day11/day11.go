package main

import (
	"bufio"
	"os"
	"strconv"
)

func readInputData(path string) [][]int {
	file, _ := os.Open(path)

	var result [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var currentNumbers []int
		row := scanner.Text()

		for _, char := range row {
			currentNumber, _ := strconv.Atoi(string(char))
			currentNumbers = append(currentNumbers, currentNumber)
		}
		result = append(result, currentNumbers)
	}
	return result
}

func main() {
	part1(readInputData("day11input.txt"))
	part2(readInputData("day11input.txt"))
}

type Point struct {
	X int
	Y int
}

func part1(inputRows [][]int) {

	sumOfFlashes := 0
	steps := 100
	for i := 0; i < steps; i++ {
		var flashes []Point
		// increase all for current step
		for currentRowIndex, currentRow := range inputRows {
			for currentColIndex, _ := range currentRow {
				if currentRow[currentColIndex] >= 9 {
					flashes = append(flashes, Point{currentRowIndex, currentColIndex})
					currentRow[currentColIndex] = 0
				} else {
					currentRow[currentColIndex]++
				}
			}
		}
		// Increace adjacents

		currentFlashes := flashes
		for {
			newFlashes := make([]Point, 0)
			for _, flash := range currentFlashes {
				// Find adjacents
				newFlashes = append(newFlashes, flashAdjacents(inputRows, flash, flashes)...)
			}
			currentFlashes = newFlashes
			flashes = append(flashes, newFlashes...)
			if len(newFlashes) == 0 {
				break
			}
		}
		sumOfFlashes += len(flashes)
	}

	println("Finished " + strconv.Itoa(steps) + " with " + strconv.Itoa(sumOfFlashes) + " number of flashes!")
}

func part2(inputRows [][]int) {

	sumOfFlashes := 0
	steps := 2000
	for i := 0; i < steps; i++ {
		var flashes []Point
		// increase all for current step
		for currentRowIndex, currentRow := range inputRows {
			for currentColIndex, _ := range currentRow {
				if currentRow[currentColIndex] >= 9 {
					flashes = append(flashes, Point{currentRowIndex, currentColIndex})
					currentRow[currentColIndex] = 0
				} else {
					currentRow[currentColIndex]++
				}
			}
		}
		// Increace adjacents

		currentFlashes := flashes
		for {
			newFlashes := make([]Point, 0)
			for _, flash := range currentFlashes {
				// Find adjacents
				newFlashes = append(newFlashes, flashAdjacents(inputRows, flash, flashes)...)
			}
			currentFlashes = newFlashes
			flashes = append(flashes, newFlashes...)
			if len(newFlashes) == 0 || len(flashes) == 100 {
				break
			}
		}
		sumOfFlashes += len(flashes)
		if len(flashes) == 100 {
			steps = i + 1
			break
		}
	}

	println("All flashing at step " + strconv.Itoa(steps) + " !")
}

func flashAdjacents(inputRows [][]int, flash Point, flashes []Point) []Point {
	newFlashes := make([]Point, 0)
	// Row above
	newFlashes = increaseAdjacent(inputRows, flash.X-1, flash.Y-1, newFlashes, flashes)
	newFlashes = increaseAdjacent(inputRows, flash.X-1, flash.Y, newFlashes, flashes)
	newFlashes = increaseAdjacent(inputRows, flash.X-1, flash.Y+1, newFlashes, flashes)
	// Left
	newFlashes = increaseAdjacent(inputRows, flash.X, flash.Y-1, newFlashes, flashes)
	// Right
	newFlashes = increaseAdjacent(inputRows, flash.X, flash.Y+1, newFlashes, flashes)
	// Row below
	newFlashes = increaseAdjacent(inputRows, flash.X+1, flash.Y-1, newFlashes, flashes)
	newFlashes = increaseAdjacent(inputRows, flash.X+1, flash.Y, newFlashes, flashes)
	newFlashes = increaseAdjacent(inputRows, flash.X+1, flash.Y+1, newFlashes, flashes)

	return newFlashes
}

func increaseAdjacent(inputRows [][]int, adjacentX int, adjacentY int, newFlashes []Point, currentFlashes []Point) []Point {
	if adjacentX >= 0 && adjacentY >= 0 && adjacentX < len(inputRows) && adjacentY < len(inputRows[0]) { // Position exists?
		if inputRows[adjacentX][adjacentY] >= 9 {
			inputRows[adjacentX][adjacentY] = 0
			if (!contains(currentFlashes, Point{adjacentX, adjacentY})) {
				newFlashes = append(newFlashes, Point{adjacentX, adjacentY})
			}
		} else if inputRows[adjacentX][adjacentY] > 0 {
			inputRows[adjacentX][adjacentY]++
		}
	}
	return newFlashes
}

func contains(s []Point, e Point) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
