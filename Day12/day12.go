package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func readInputData(path string) map[string]Point {
	file, _ := os.Open(path)

	result := make(map[string]Point)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//var currentNumbers []int
		row := scanner.Text()
		splittedStrings := strings.Split(row, "-")

		startPoint := splittedStrings[0]
		endPoint := splittedStrings[1]
		_, keyExistsStart := result[startPoint]
		_, keyExistsEnd := result[endPoint]
		var currentStartPoint Point
		var currentEndPoint Point

		if keyExistsStart {
			currentStartPoint = result[startPoint]
		} else {
			currentStartPoint = Point{Name: startPoint, IsUppercase: IsUpper(startPoint)}
		}

		if keyExistsEnd {
			currentEndPoint = result[endPoint]
		} else {
			currentEndPoint = Point{Name: endPoint, IsUppercase: IsUpper(endPoint)}
		}

		if !contains(currentStartPoint.Neighbours, endPoint) {
			currentStartPoint.Neighbours = append(currentStartPoint.Neighbours, endPoint)
		}
		if !contains(currentEndPoint.Neighbours, startPoint) {
			currentEndPoint.Neighbours = append(currentEndPoint.Neighbours, startPoint)
		}

		result[startPoint] = currentStartPoint
		result[endPoint] = currentEndPoint
	}
	return result
}

func main() {
	part1(readInputData("day12input.txt"))

}

type Point struct {
	Name        string
	IsUppercase bool
	Neighbours  []string
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func contains(input []string, item string) bool {
	for _, a := range input {
		if a == item {
			return true
		}
	}
	return false
}

func part1(input map[string]Point) {
	startPoint := input["start"]
	var paths []string
	for _, neighbour := range startPoint.Neighbours {

		for {
			var currentPath []string
			currentPath = append(currentPath, startPoint.Name)
			currentPath = append(currentPath, neighbour)

			var digMore bool
			currentPath, digMore, paths = VisitNeighbours(input, input[neighbour], currentPath, paths)

			if !digMore {
				break
			}
		}
	}

	for _, path := range paths {
		println(path)
	}

	println("Number of paths found is: " + strconv.Itoa(len(paths)))
}

func VisitNeighbours(input map[string]Point, currentPoint Point, currentPath []string, paths []string) ([]string, bool, []string) {

	digMore := false

	for _, neighbour := range currentPoint.Neighbours {
		currentworkingPath := currentPath

		if (IsUpper(neighbour) || !contains(currentworkingPath, neighbour)) && !Visited(currentworkingPath, neighbour, paths) {
			digMore = true
			currentworkingPath = append(currentworkingPath, neighbour)
			if neighbour == "end" {
				currentPathStr := ""

				for _, node := range currentworkingPath {
					currentPathStr += node + ","
				}
				paths = append(paths, currentPathStr[:len(currentPathStr)-1])

				continue
			}
			currentworkingPath, digMore, paths = VisitNeighbours(input, input[neighbour], currentworkingPath, paths)

			if currentworkingPath[len(currentPath)-1] == "end" {
				currentPathStr := ""

				for _, node := range currentworkingPath {
					currentPathStr += node + ","
				}
				paths = append(paths, currentPathStr[:len(currentPathStr)-1])

				continue
			}
		}
	}
	return currentPath, digMore, paths
}

func Visited(currentPath []string, nextHop string, paths []string) bool {
	// Check if current path matches one alredy created
	var currentPathStr string
	for _, node := range currentPath {
		currentPathStr += node + ","
	}
	currentPathStr += nextHop

	for _, pathStr := range paths {
		if len(pathStr) >= len(currentPathStr) && currentPathStr == pathStr[:len(currentPathStr)] {
			return true
		}
	}
	return false
}
