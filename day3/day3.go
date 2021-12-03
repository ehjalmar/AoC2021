package main

import (
	"bufio"
	"os"
	"strconv"
)

type DiagnosticReportRow struct {
	Values []int
}

func readInputData(path string) ([]DiagnosticReportRow, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var commands []DiagnosticReportRow
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		var newRow DiagnosticReportRow
		for _, character := range data {
			newValue, _ := strconv.Atoi(string(character))
			newRow.Values = append(newRow.Values, newValue)
		}
		commands = append(commands, newRow)
	}
	return commands, scanner.Err()
}

func main() {
	inputData, _ := readInputData("day3input.txt")

	executePart1(inputData)
	//executePart2(commands)
}

func executePart1(inputData []DiagnosticReportRow) {
	var gamma string
	var epsilon string

	for i := 0; i < len(inputData[0].Values); i++ {
		zeros := 0
		ones := 0
		for _, row := range inputData {
			if row.Values[i] == 0 {
				zeros++
			} else if row.Values[i] == 1 {
				ones++
			}
		}
		if ones > zeros {
			// gamma.Values[i] = 1
			// epsilon.Values[i] = 0
			gamma = gamma + strconv.Itoa(1)
			epsilon = epsilon + strconv.Itoa(0)
		} else {
			// gamma.Values[i] = 0
			// epsilon.Values[i] = 1
			gamma = gamma + strconv.Itoa(0)
			epsilon = epsilon + strconv.Itoa(1)
		}
	}
	println(gamma)
	println(epsilon)
	gammaDecimal, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonDecimal, _ := strconv.ParseInt(epsilon, 2, 64)
	println(gammaDecimal * epsilonDecimal)
}
