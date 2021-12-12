package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func readInputData(path string) []string {
	file, _ := os.Open(path)

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		result = append(result, row)
	}
	return result
}

func main() {
	incompleteLines := part1(readInputData("day10input.txt"))
	part2(incompleteLines)
}

func part1(input []string) []string {
	var incompleteRows []string
	charScores := make([]int, 4)

	for _, row := range input {
		var rowChars []string
		corruptedRow := false
		for _, char := range row {
			currentChar := string(char)
			if currentChar == "{" || currentChar == "[" || currentChar == "(" || currentChar == "<" {
				rowChars = append(rowChars, currentChar)
			} else {
				lastChar := rowChars[len(rowChars)-1]
				if lastChar == "{" {
					if currentChar == "}" {
						rowChars = rowChars[:len(rowChars)-1]
					} else {
						corruptedRow = true
						setCharScore(currentChar, charScores)
						break
					}
				} else if lastChar == "[" {
					if currentChar == "]" {
						rowChars = rowChars[:len(rowChars)-1]
					} else {
						corruptedRow = true
						setCharScore(currentChar, charScores)
						break
					}
				} else if lastChar == "(" {
					if currentChar == ")" {
						rowChars = rowChars[:len(rowChars)-1]
					} else {
						corruptedRow = true
						setCharScore(currentChar, charScores)
						break
					}
				} else if lastChar == "<" {
					if currentChar == ">" {
						rowChars = rowChars[:len(rowChars)-1]
					} else {
						corruptedRow = true
						setCharScore(currentChar, charScores)
						break
					}
				}
			}
		}
		if !corruptedRow {
			incompleteRows = append(incompleteRows, row)
		}
	}
	sum := 0
	for _, v := range charScores {
		sum += v
	}

	println("Total syntx error score is: " + strconv.Itoa(sum))
	return incompleteRows
}

func setCharScore(char string, charScores []int) {
	switch char {
	case ")":
		charScores[0] += 3
	case "]":
		charScores[1] += 57
	case "}":
		charScores[2] += 1197
	case ">":
		charScores[3] += 25137
	}
}

func part2(input []string) {
	var rowScores []int
	for _, row := range input {
		completionString := ""

		for i := 0; i < len(row); i++ {
			currentChar := string(row[i])

			switch currentChar {
			case "(":
				completionString += GetMissingClosings(i, row, "(", ")")
			case "[":
				completionString += GetMissingClosings(i, row, "[", "]")

			case "{":
				completionString += GetMissingClosings(i, row, "{", "}")
			case "<":
				completionString += GetMissingClosings(i, row, "<", ">")
			}
		}

		currentRowScore := 0
		for i := len(completionString) - 1; i >= 0; i-- {
			currentClosingChar := string(completionString[i])
			var charPoints int
			switch currentClosingChar {
			case ")":
				charPoints = 1
			case "]":
				charPoints = 2
			case "}":
				charPoints = 3
			case ">":
				charPoints = 4
			}

			currentRowScore = getScore(currentRowScore, charPoints)
		}

		rowScores = append(rowScores, currentRowScore)
	}
	sort.Ints(rowScores)

	sum := rowScores[len(rowScores)/2]

	println("Middle score is: " + strconv.Itoa(sum))
}

func GetMissingClosings(currentRowIndex int, currentRow string, startChar string, closingChar string) string {
	openingsOfType := 1
	foundClosing := false
	completionString := ""
	for foo := currentRowIndex + 1; foo < len(currentRow); foo++ {
		if string(currentRow[foo]) == closingChar {
			if openingsOfType == 1 {
				foundClosing = true
			} else {
				openingsOfType--
			}
		} else if string(currentRow[foo]) == startChar {
			openingsOfType++
		}
	}
	if !foundClosing {
		completionString += closingChar
	}
	return completionString
}

func getScore(currentScore int, charPoints int) int {
	return (currentScore * 5) + charPoints
}
