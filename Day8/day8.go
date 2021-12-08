package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type SegmentDisplay struct {
	Input  []string
	Output []string
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

func readInputData(path string) []SegmentDisplay {
	file, _ := os.Open(path)
	var currentRow SegmentDisplay
	var displays []SegmentDisplay

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splittedStrings := strings.Split(scanner.Text(), "|")
		var input []string
		inputRange := deleteEmpty(strings.Split(splittedStrings[0], " "))
		for _, currentValue := range inputRange {
			input = append(input, currentValue)
		}
		var output []string
		outputRange := deleteEmpty(strings.Split(splittedStrings[1], " "))
		for _, currentValue := range outputRange {
			output = append(output, currentValue)
		}
		currentRow = SegmentDisplay{input, output}
		displays = append(displays, currentRow)
	}

	return displays
}

func main() {
	part1(readInputData("day8input.txt"))
	part2(readInputData("day8input.txt"))
}

func part1(input []SegmentDisplay) {

	result := make(map[int]int)

	for _, currentValue := range input {
		for _, currentOutput := range currentValue.Output {
			switch len(currentOutput) {
			case 2:
				result[1]++
			case 4:
				result[4]++
			case 3:
				result[7]++
			case 7:
				result[8]++
			}
		}
	}
	sum := 0
	for key, currentResult := range result {
		println(strconv.Itoa(key) + " " + strconv.Itoa(currentResult))
		sum += currentResult
	}
	println("1, 4, 7, or 8 appears " + strconv.Itoa(sum) + " times.")
}

func part2(input []SegmentDisplay) {

	sum := 0

	for _, currentValue := range input {

		signalMapping := make(map[string]string)
		facit := make(map[int]string)

		for _, currentInput := range currentValue.Input {
			switch len(currentInput) {
			case 2:
				signalMapping[currentInput] = "1"
				facit[1] = currentInput
			case 4:
				signalMapping[currentInput] = "4"
				facit[4] = currentInput
			case 3:
				signalMapping[currentInput] = "7"
				facit[7] = currentInput
			case 7:
				signalMapping[currentInput] = "8"
				facit[8] = currentInput
			}
		}
		for _, currentInput := range currentValue.Input {
			if len(currentInput) == 6 || len(currentInput) == 5 {
				existsIn1 := 0
				existsIn4 := 0
				existsIn7 := 0
				existsIn8 := 0

				for _, char := range currentInput {

					if strings.Contains(facit[1], string(char)) {
						existsIn1++
					}
					if strings.Contains(facit[4], string(char)) {
						existsIn4++
					}
					if strings.Contains(facit[7], string(char)) {
						existsIn7++
					}
					if strings.Contains(facit[8], string(char)) {
						existsIn8++
					}
				}

				if existsIn4 == 3 && existsIn1 == 2 && len(currentInput) == 6 { // 3 == 4 && 2 == 1 blir 0
					signalMapping[currentInput] = "0"
				} else if existsIn8 == 6 && existsIn1 == 1 && len(currentInput) == 6 { // 1 = 1 && 6 == 8 blir 6
					signalMapping[currentInput] = "6"
				} else if existsIn8 == 5 && existsIn1 == 2 && len(currentInput) == 5 { // 5 = 8 & 2 == 1 blir 3
					signalMapping[currentInput] = "3"
				} else if existsIn8 == 5 && existsIn4 == 2 && len(currentInput) == 5 { // 5 = 8 && 1 == 1 blir 2
					signalMapping[currentInput] = "2"
				} else if existsIn4 == 4 && existsIn7 == 3 && len(currentInput) == 6 { // 4 = 4 && 3 == 7 blir 9
					signalMapping[currentInput] = "9"
				} else if existsIn4 == 3 && existsIn1 == 1 && len(currentInput) == 5 { // 3 = 4 && 1 == 1 blir 5
					signalMapping[currentInput] = "5"
				}
			}
		}

		currentDigits := ""

		for _, currentOutput := range currentValue.Output {
			//if same length and all chars exists
			for currentSignalMapping, currentDigit := range signalMapping {
				currentLength := len(currentOutput)
				exists := 0
				if currentLength == len(currentSignalMapping) {
					for _, char := range currentSignalMapping {
						if strings.Contains(currentOutput, string(char)) {
							exists++
						}
					}
					if exists == currentLength {
						currentDigits += currentDigit
						break
					}
				}

			}
		}
		currentSum, _ := strconv.Atoi(currentDigits)
		sum += currentSum
	}

	println("Sum of outputvaules: " + strconv.Itoa(sum))
}
