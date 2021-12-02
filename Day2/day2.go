package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	CommandName string
	Distance    int
}

func main() {
	commands, _ := readInputNumbers("inputday2.txt")

	executePart1(commands)
	executePart2(commands)
}

func readInputNumbers(path string) ([]Command, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var commands []Command
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputValues := strings.Split(scanner.Text(), " ")
		val1 := inputValues[0]
		val2, _ := strconv.Atoi(inputValues[1])
		commands = append(commands, Command{val1, val2})
	}
	return commands, scanner.Err()
}

func executePart1(commands []Command) {
	x := 0
	y := 0

	for _, command := range commands {
		if command.CommandName == "forward" {
			x = x + command.Distance
		} else if command.CommandName == "down" {
			y = y + command.Distance
		} else if command.CommandName == "up" {
			y = y - command.Distance
		}
	}
	println(x)
	println(y)
	println(x * y)
}

func executePart2(commands []Command) {
	x := 0
	y := 0
	aim := 0

	for _, command := range commands {
		if command.CommandName == "forward" {
			x = x + command.Distance
			y = y + (aim * command.Distance)
		} else if command.CommandName == "down" {
			aim = aim + command.Distance
		} else if command.CommandName == "up" {
			aim = aim - command.Distance
		}
	}
	println(x)
	println(y)
	println(x * y)
}
