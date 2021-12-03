package main

import (
	"bufio"
	"fmt"
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
	executePart2(inputData)
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
			gamma = gamma + strconv.Itoa(1)
			epsilon = epsilon + strconv.Itoa(0)
		} else if ones == zeros {
			gamma = gamma + strconv.Itoa(1)
			epsilon = epsilon + strconv.Itoa(0)
		} else {
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

func mostCommonForPosition(inputData []DiagnosticReportRow, position int) int {

	zeros := 0
	ones := 0
	for _, row := range inputData {
		if row.Values[position] == 0 {
			zeros++
		} else if row.Values[position] == 1 {
			ones++
		}
	}
	if ones >= zeros {
		return 1
	} else {
		return 0
	}
}

func leastCommonForPosition(inputData []DiagnosticReportRow, position int) int {

	zeros := 0
	ones := 0
	for _, row := range inputData {
		if row.Values[position] == 0 {
			zeros++
		} else if row.Values[position] == 1 {
			ones++
		}
	}
	if ones >= zeros {
		return 0
	} else {
		return 1
	}
}

func executePart2(inputData []DiagnosticReportRow) {

	inputDataEpsilon := inputData
	var gamma string
	var epsilon string

	for pos := 0; pos < len(inputData[0].Values); pos++ {
		var dataToKeepGamma []DiagnosticReportRow
		mostCommon := mostCommonForPosition(inputData, pos)

		for i := 0; i < len(inputData); i++ {
			if inputData[i].Values[pos] == mostCommon {
				dataToKeepGamma = append(dataToKeepGamma, inputData[i])
			}
		}
		inputData = dataToKeepGamma
		if len(dataToKeepGamma) == 1 {
			break
		}
	}
	println("Gamma:")
	for _, v := range inputData[0].Values {
		print(v)
		gamma = gamma + strconv.Itoa(v)
	}

	for pos := 0; pos < len(inputDataEpsilon[0].Values); pos++ {
		var dataToKeepEpsilon []DiagnosticReportRow
		lestCommon := mostCommonForPosition(inputDataEpsilon, pos)

		for i := 0; i < len(inputDataEpsilon); i++ {
			if inputDataEpsilon[i].Values[pos] != lestCommon {
				dataToKeepEpsilon = append(dataToKeepEpsilon, inputDataEpsilon[i])
			}
		}
		inputDataEpsilon = dataToKeepEpsilon
		if len(dataToKeepEpsilon) == 1 {
			break
		}
	}
	println("\nEpsilon:")

	for _, v := range inputDataEpsilon[0].Values {
		print(v)
		epsilon = epsilon + strconv.Itoa(v)
	}

	gammaDecimal, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonDecimal, _ := strconv.ParseInt(epsilon, 2, 64)
	println("\n" + fmt.Sprint(gammaDecimal*epsilonDecimal))
}
