package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

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
	lowPoints := part1(readInputData("day9input.txt"))
	part2(readInputData("day9input.txt"), lowPoints)
}

func part1(input [][]int) []Point {
	var lowPoints []Point
	var sum int
	for rowIndex, currentRow := range input {
		for columnIndex, currentNumber := range currentRow {
			lessThanSouth := false
			lessThanNorth := false
			lessThanWest := false
			lessThanEast := false
			if rowIndex == (len(input)-1) || currentNumber < input[rowIndex+1][columnIndex] {
				lessThanSouth = true
			}
			if columnIndex == (len(currentRow)-1) || currentNumber < currentRow[columnIndex+1] {
				lessThanEast = true
			}
			if columnIndex == 0 || currentNumber < currentRow[columnIndex-1] {
				lessThanWest = true
			}
			if rowIndex == 0 || currentNumber < input[rowIndex-1][columnIndex] {
				lessThanNorth = true
			}

			if lessThanEast && lessThanNorth && lessThanSouth && lessThanWest {
				//println(strconv.Itoa(currentNumber))
				sum += currentNumber + 1
				lowPoints = append(lowPoints, Point{rowIndex, columnIndex, currentNumber})
			}
		}
	}
	println("Sum of risklevel is: " + strconv.Itoa(sum))
	return lowPoints
}

type Point struct {
	Row    int
	Column int
	Value  int
}

func part2(input [][]int, lowPoints []Point) {

	var basins []int
	for _, currentLowPoint := range lowPoints {

		var visited []Point
		basinSum := 1
		basinSum, _ = findBasinMembers(currentLowPoint, basinSum, input, visited)

		basins = append(basins, basinSum)
		// println("Current lowPoint: " + strconv.Itoa(currentLowPoint.Row) + " " + strconv.Itoa(currentLowPoint.Column))
		// println("Current value: " + strconv.Itoa(input[currentLowPoint.Row][currentLowPoint.Column]))
	}

	sort.Slice(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})

	sum := basins[0] * basins[1] * basins[2]

	println("Three largest basins multiplied is: " + strconv.Itoa(sum))
}

func findBasinMembers(currentLowPoint Point, basinSum int, input [][]int, visited []Point) (int, []Point) {
	// Look up
	if currentLowPoint.Row > 0 && input[currentLowPoint.Row-1][currentLowPoint.Column] > currentLowPoint.Value && input[currentLowPoint.Row-1][currentLowPoint.Column] < 9 && isVisited(currentLowPoint.Row-1, currentLowPoint.Column, visited) == false {
		basinSum++
		visited = append(visited, Point{currentLowPoint.Row - 1, currentLowPoint.Column, currentLowPoint.Value + 1})
		basinSum, visited = findBasinMembers(Point{currentLowPoint.Row - 1, currentLowPoint.Column, input[currentLowPoint.Row-1][currentLowPoint.Column]}, basinSum, input, visited)
	}
	// Look east
	if currentLowPoint.Column < (len(input[0])-1) && input[currentLowPoint.Row][currentLowPoint.Column+1] > currentLowPoint.Value && input[currentLowPoint.Row][currentLowPoint.Column+1] < 9 && isVisited(currentLowPoint.Row, currentLowPoint.Column+1, visited) == false {
		basinSum++
		visited = append(visited, Point{currentLowPoint.Row, currentLowPoint.Column + 1, currentLowPoint.Value + 1})
		basinSum, visited = findBasinMembers(Point{currentLowPoint.Row, currentLowPoint.Column + 1, input[currentLowPoint.Row][currentLowPoint.Column+1]}, basinSum, input, visited)
	}
	// Look south
	if currentLowPoint.Row < (len(input)-1) && input[currentLowPoint.Row+1][currentLowPoint.Column] > currentLowPoint.Value && input[currentLowPoint.Row+1][currentLowPoint.Column] < 9 && isVisited(currentLowPoint.Row+1, currentLowPoint.Column, visited) == false {
		basinSum++
		visited = append(visited, Point{currentLowPoint.Row + 1, currentLowPoint.Column, currentLowPoint.Value + 1})
		basinSum, visited = findBasinMembers(Point{currentLowPoint.Row + 1, currentLowPoint.Column, input[currentLowPoint.Row+1][currentLowPoint.Column]}, basinSum, input, visited)
	}
	// Look west
	if currentLowPoint.Column > 0 && input[currentLowPoint.Row][currentLowPoint.Column-1] > currentLowPoint.Value && input[currentLowPoint.Row][currentLowPoint.Column-1] < 9 && isVisited(currentLowPoint.Row, currentLowPoint.Column-1, visited) == false {
		basinSum++
		visited = append(visited, Point{currentLowPoint.Row, currentLowPoint.Column - 1, currentLowPoint.Value + 1})
		basinSum, visited = findBasinMembers(Point{currentLowPoint.Row, currentLowPoint.Column - 1, input[currentLowPoint.Row][currentLowPoint.Column-1]}, basinSum, input, visited)
	}
	return basinSum, visited
}

func isVisited(row int, col int, visited []Point) bool {
	for _, v := range visited {
		if v.Row == row && v.Column == col {
			return true
		}
	}
	return false
}
