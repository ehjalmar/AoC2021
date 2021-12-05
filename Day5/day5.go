package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	Start DataPoint
	End   DataPoint
}

type DataPoint struct {
	X int
	Y int
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func readInputData(path string) ([]Line, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rows []Line
	var currentLine Line
	var currentDataPoint DataPoint

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()

		splittedData := strings.Split(data, " -> ")
		for i := 0; i < 2; i++ {
			cords := strings.Split(splittedData[i], ",")
			if i == 0 {
				currentDataPoint.X, _ = strconv.Atoi(cords[i])
				currentDataPoint.Y, _ = strconv.Atoi(cords[i+1])
				currentLine.Start = currentDataPoint
			} else if i == 1 {
				currentDataPoint.X, _ = strconv.Atoi(cords[i-1])
				currentDataPoint.Y, _ = strconv.Atoi(cords[i])
				currentLine.End = currentDataPoint
			}
		}
		rows = append(rows, currentLine)
	}
	return rows, scanner.Err()
}

func main() {
	lines, _ := readInputData("day5input.txt")
	calculateOverlap1(lines)
	calculateOverlap2(lines)
}

func calculateOverlap1(lines []Line) {

	visitedDataPoints := make(map[DataPoint]int)
	//var currentPoint string
	var currentPointD DataPoint

	for _, line := range lines {
		if line.Start.X == line.End.X { // Vertical or horizontal?
			if line.Start.Y < line.End.Y { // Going up or down?
				diff := line.End.Y - line.Start.Y
				for i := 0; i <= diff; i++ {
					currentPointD = DataPoint{line.Start.X, line.Start.Y + i}
					visitedDataPoints[currentPointD] += 1
				}
			} else if line.Start.Y > line.End.Y {
				diff := line.Start.Y - line.End.Y
				for i := 0; i <= diff; i++ {
					currentPointD = DataPoint{line.Start.X, line.Start.Y - i}
					visitedDataPoints[currentPointD] += 1
				}
			}
		} else if line.Start.Y == line.End.Y {
			if line.Start.X < line.End.X {
				diff := line.End.X - line.Start.X
				for i := 0; i <= diff; i++ {
					currentPointD = DataPoint{line.Start.X + i, line.Start.Y}
					visitedDataPoints[currentPointD] += 1
				}
			} else if line.Start.X > line.End.X {
				diff := line.Start.X - line.End.X
				for i := 0; i <= diff; i++ {
					currentPointD = DataPoint{line.Start.X - i, line.Start.Y}
					visitedDataPoints[currentPointD] += 1
				}
			}
		}
	}
	pointsWithOverlap := 0
	for _, visitedDataPoint := range visitedDataPoints {
		if visitedDataPoint > 1 {
			pointsWithOverlap++
		}
	}
	println("Points with more than two visists: " + strconv.Itoa(pointsWithOverlap))
}

func calculateOverlap2(lines []Line) {

	visitedDataPoints := make(map[DataPoint]int)
	//var currentPoint string
	var currentPointD DataPoint

	for _, line := range lines {
		if line.Start.X == line.End.X { // Vertical, horizontal or diagonal?
			if line.Start.Y < line.End.Y { // Going up or down?
				diff := line.End.Y - line.Start.Y
				for i := 0; i <= diff; i++ {
					currentPointD = DataPoint{line.Start.X, line.Start.Y + i}
					visitedDataPoints[currentPointD] += 1
				}
			} else if line.Start.Y > line.End.Y {
				diff := line.Start.Y - line.End.Y
				for i := 0; i <= diff; i++ {
					currentPointD = DataPoint{line.Start.X, line.Start.Y - i}
					visitedDataPoints[currentPointD] += 1
				}
			}
		} else if line.Start.Y == line.End.Y {
			if line.Start.X < line.End.X {
				diff := line.End.X - line.Start.X
				for i := 0; i <= diff; i++ {
					currentPointD = DataPoint{line.Start.X + i, line.Start.Y}
					visitedDataPoints[currentPointD] += 1
				}
			} else if line.Start.X > line.End.X {
				diff := line.Start.X - line.End.X
				for i := 0; i <= diff; i++ {
					currentPointD = DataPoint{line.Start.X - i, line.Start.Y}
					visitedDataPoints[currentPointD] += 1
				}
			}
		} else {
			if line.Start.X < line.End.X && line.Start.Y < line.End.Y { // X+ Y+
				diff := line.End.X - line.Start.X
				for i := 0; i <= diff; i++ {
					currentPointD = DataPoint{line.Start.X + i, line.Start.Y + i}
					visitedDataPoints[currentPointD] += 1
				}
			} else if line.Start.X > line.End.X && line.Start.Y < line.End.Y { // X- Y+
				diff := line.Start.X - line.End.X
				for i := 0; i <= diff; i++ {
					currentPointD = DataPoint{line.Start.X - i, line.Start.Y + i}
					visitedDataPoints[currentPointD] += 1
				}
			} else if line.Start.X < line.End.X && line.Start.Y > line.End.Y { // X+ Y-
				diff := line.End.X - line.Start.X
				for i := 0; i <= diff; i++ {
					currentPointD = DataPoint{line.Start.X + i, line.Start.Y - i}
					visitedDataPoints[currentPointD] += 1
				}
			} else if line.Start.X > line.End.X && line.Start.Y > line.End.Y { // X- Y-
				diff := line.Start.X - line.End.X
				for i := 0; i <= diff; i++ {
					currentPointD = DataPoint{line.Start.X - i, line.Start.Y - i}
					visitedDataPoints[currentPointD] += 1
				}
			}
		}
	}
	pointsWithOverlap := 0
	for _, visitedDataPoint := range visitedDataPoints {
		if visitedDataPoint > 1 {
			pointsWithOverlap++
		}
	}
	println("Points with more than two visists: " + strconv.Itoa(pointsWithOverlap))
}
