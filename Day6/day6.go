package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func readInputData(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splittedStrings := strings.Split(scanner.Text(), ",")
		for _, i := range splittedStrings {
			newValue, _ := strconv.Atoi(i)
			j := int(newValue)

			lines = append(lines, j)
		}
	}
	return lines, scanner.Err()
}

func main() {
	lines, _ := readInputData("day6input.txt")
	countFish(lines, 80)
	countFish(lines, 256)
}

func countFish(fish []int, days int) {

	dayFish := make(map[int]int)
	fishMap := make(map[int]int)

	for i := 0; i <= 8; i++ {
		dayFish[i] = 0
		fishMap[i] = 0
	}

	for i := 0; i < len(fish); i++ {
		fishMap[fish[i]]++
		dayFish[fish[i]]++
	}

	for dayNumber := 0; dayNumber < days; dayNumber++ {
		for key, value := range dayFish {

			if key == 0 {
				fishMap[6] += value
				fishMap[8] += value
				fishMap[0] -= value
			} else {
				fishMap[key-1] += value
				fishMap[key] -= value
			}
		}
		for k, v := range fishMap {
			dayFish[k] = v
		}

		//println("Day: " + strconv.Itoa(dayNumber))
		//printMemUsage()
	}
	printMemUsage()
	result := 0
	for _, v := range dayFish {
		result += v
	}

	println("-----------------Number of fish is: " + strconv.Itoa(result) + " after " + strconv.Itoa(days) + " days.-----------------")
}

func bToMbyte(b uint64) float64 {
	return float64(b) / float64(1024) / float64(1024)
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	//   https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %0.2f MiB", bToMbyte(m.Alloc))
	fmt.Printf("\tTotalAlloc = %0.2f MiB", bToMbyte(m.TotalAlloc))
	fmt.Printf("\tSys = %0.2f MiB", bToMbyte(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
