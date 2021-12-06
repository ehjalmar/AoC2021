package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readInputData(path string) map[int]int {
	file, _ := os.Open(path)
	fishMap := make(map[int]int)

	for i := 0; i <= 8; i++ {
		fishMap[i] = 0
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splittedStrings := strings.Split(scanner.Text(), ",")
		for _, i := range splittedStrings {
			newValue, _ := strconv.Atoi(i)
			fishMap[newValue]++
		}
	}
	return fishMap
}

func main() {
	countFish(readInputData("day6input.txt"), 80)
	countFish(readInputData("day6input.txt"), 256)
}

func countFish(fish map[int]int, days int) {
	dayFish := make(map[int]int)
	for k, v := range fish {
		dayFish[k] = v
	}

	for dayNumber := 0; dayNumber < days; dayNumber++ {
		for key, value := range dayFish {

			if key == 0 {
				fish[6] += value
				fish[8] += value
				fish[0] -= value
			} else {
				fish[key-1] += value
				fish[key] -= value
			}
		}
		for k, v := range fish {
			dayFish[k] = v
		}
	}
	result := 0
	for _, v := range dayFish {
		result += v
	}

	println("-----------------Number of fish is: " + strconv.Itoa(result) + " after " + strconv.Itoa(days) + " days.-----------------")
}
