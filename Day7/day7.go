package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInputData(path string) map[int]int {
	file, _ := os.Open(path)
	countForPositions := make(map[int]int)
	var numbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splittedStrings := strings.Split(scanner.Text(), ",")
		for _, i := range splittedStrings {
			newValue, _ := strconv.Atoi(i)
			numbers = append(numbers, newValue)
		}
	}
	sort.Ints(numbers)
	largetsNumber := numbers[len(numbers)-1]

	for i := 0; i <= largetsNumber; i++ {
		countOnCurrent := 0
		for _, number := range numbers {
			if number == i {
				countOnCurrent++
			}
		}
		countForPositions[i] = countOnCurrent
	}

	return countForPositions
}

func main() {
	FindChepestPosition(readInputData("day7input.txt"))
}

func FindChepestPosition(countForPositions map[int]int) {
	costForPosition := make(map[int]int)

	for i := 0; i < len(countForPositions); i++ {
		var cost int
		for currentNumber, currentCount := range countForPositions {
			distance := Abs(Abs(currentNumber) - i)
			cost = distance * currentCount
			costForPosition[i] += cost
		}
		println(strconv.Itoa(i) + " has cost: " + strconv.Itoa(costForPosition[i]))
	}

	positions := make([]int, 0, len(costForPosition))
	for position := range costForPosition {
		positions = append(positions, position)
	}

	sort.Slice(positions, func(i, j int) bool {
		return costForPosition[positions[i]] > costForPosition[positions[j]]
	})

	for _, position := range positions {
		fmt.Printf("%-7v %v\n", position, costForPosition[position])
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
