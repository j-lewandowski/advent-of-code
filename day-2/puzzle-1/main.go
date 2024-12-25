package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func appendReports(reports [][]int, report string) [][]int {
	levels := strings.Fields(report)
	levelsList := make([]int, 0)

	for _, level := range(levels){
		value, err := strconv.Atoi(level)

		if(err != nil) {
			panic("Not a number")
		}

		levelsList = append(levelsList, value)
	}

	return append(reports, levelsList)
}


func isSafe(report []int) bool {

	isIncreasing, isDecreasing := true, true
	
	for i := range(len(report)-1){
		if (report[i] >= report[i+1]) {
			isDecreasing = false;
		}

		if (report[i] <= report[i+1]) {
			isIncreasing = false;
		}

		if (math.Abs(float64(report[i] - report[i+1])) > 3) {
			return false
		}
	}

	return isDecreasing || isIncreasing
}


func main() {
	file, err := os.Open("input.txt")
	if (err != nil) {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	reports := make([][]int, 0)

	for scanner.Scan() {
		reports = appendReports(reports ,scanner.Text())
	}

	counter := 0

	for _, report := range(reports){
		if (isSafe(report)) {
			counter++
		}
	}

	fmt.Println(counter)
}
