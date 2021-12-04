package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BingoBoard struct {
	Rows               [][]int
	Columns            map[int][]int
	MatchNumbersForRow map[int]int
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

func readInputData(path string) ([]BingoBoard, []int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	rowNumber := 0
	var rows []BingoBoard
	var currentBoard BingoBoard
	var numbers []int
	var columns map[int][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()

		if rowNumber == 0 {
			splittedStrings := strings.Split(data, ",")

			for _, i := range splittedStrings {
				j, err := strconv.Atoi(i)
				if err != nil {
					panic(err)
				}
				numbers = append(numbers, j)
			}
		} else {

			if data == "" {
				currentBoard = BingoBoard{}
				currentBoard.MatchNumbersForRow = make(map[int]int)
				columns = make(map[int][]int)
			} else {
				splittedStrings := strings.Split(strings.TrimSpace(data), " ")
				splittedStrings = deleteEmpty(splittedStrings)
				var currentBoardRow []int

				for index, i := range splittedStrings {
					j, err := strconv.Atoi(i)
					if err != nil {
						panic(err)
					}
					currentBoardRow = append(currentBoardRow, j)
					columns[index] = append(columns[index], j)
				}

				currentBoard.Rows = append(currentBoard.Rows, currentBoardRow)
			}
			if len(currentBoard.Rows) == 5 {
				for i := 0; i < len(columns); i++ {
					currentBoard.Rows = append(currentBoard.Rows, columns[i])
				}
				rows = append(rows, currentBoard)
			}
		}
		rowNumber++
	}
	return rows, numbers, scanner.Err()
}

func main() {
	inputData, numbers, _ := readInputData("day4input.txt")
	playBingo(inputData, numbers, true)
	inputData, numbers, _ = readInputData("day4input.txt")
	playBingo(inputData, numbers, false)
}

func playBingo(inputData []BingoBoard, numbers []int, stopOnFirstWinner bool) {
	var markedNumbers = make(map[int]bool)
	winningBoards := make(map[int]bool)

	for _, currentNumber := range numbers {
		markedNumbers[currentNumber] = true
		for currentBoardIndex, currentBoard := range inputData {
			for rowIndex, currentBoardRow := range currentBoard.Rows {
				for _, currentBoardNumber := range currentBoardRow {
					if currentNumber == currentBoardNumber {
						currentBoard.MatchNumbersForRow[rowIndex] += 1
						if currentBoard.MatchNumbersForRow[rowIndex] == 5 {
							winningBoards[currentBoardIndex] = true
							if stopOnFirstWinner {
								println("MATCH ON ROW/COLUMN NUMBER(zero index adjusted): " + strconv.Itoa(rowIndex+1) + " for board number: " + strconv.Itoa(currentBoardIndex+1))
								println("Winner row: ")
								fmt.Printf("%v\n", currentBoardRow)

								sum := getSumOfUnmarkedNumbers(currentBoard, markedNumbers)
								println("Final score for first winning board is: " + strconv.Itoa(sum*currentBoardNumber))
								return
							} else {
								if len(winningBoards) == len(inputData) {
									println("Last board to win is: " + strconv.Itoa(currentBoardIndex+1))
									sum := getSumOfUnmarkedNumbers(currentBoard, markedNumbers)
									println("Final score for last winning board is: " + strconv.Itoa(sum*currentBoardNumber))
									return
								}
							}
						}
					}
				}
			}
		}
	}
}

func getSumOfUnmarkedNumbers(currentBoard BingoBoard, markedNumbers map[int]bool) int {
	sum := 0
	for rowIndex, currentBoardRow := range currentBoard.Rows {
		for _, currentBoardNumber := range currentBoardRow {
			if !markedNumbers[currentBoardNumber] {
				sum += currentBoardNumber
			}
		}
		if rowIndex == 4 { // Don´t count columns
			return sum
		}
	}
	return sum
}
