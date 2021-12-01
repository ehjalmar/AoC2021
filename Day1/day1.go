package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines, _ := readInputNumbers("inputday1.txt")

	executePart1(lines)
	executePart2(lines)
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readInputNumbers(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputNumber, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, inputNumber)
	}
	return lines, scanner.Err()
}

func executePart1(lines []int) {
	depthIncreased := 0

	for i, line := range lines {
		if i > 0 && line > lines[i-1] {
			depthIncreased++
		}
	}
	fmt.Println(fmt.Sprint("Part 1 - Depth increased number of times: ", depthIncreased))
}

func executePart2(lines []int) {
	depthIncreased := 0

	for i, _ := range lines {
		if i < len(lines)-3 { // Fix end
			left := lines[i] + lines[i+1] + lines[i+2]
			right := lines[i+1] + lines[i+2] + lines[i+3]
			if right > left {
				depthIncreased++
			}
		}
	}
	fmt.Println(fmt.Sprint("Part2 - Depth increased number of times: ", depthIncreased))
}
