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

	for _, level := range levels {
		value, err := strconv.Atoi(level)
		if err != nil {
			panic("Not a number")
		}

		levelsList = append(levelsList, value)
	}

	return append(reports, levelsList)
}

func isSafe(report []int) bool {
	checkSafe := func(report []int) bool {
			isSubtle := true
			isDecreasing, isIncreasing := true, true

			for i := range report[:len(report)-1] {
					if report[i] >= report[i+1] {
							isDecreasing = false
					}

					if report[i] <= report[i+1] {
							isIncreasing = false
					}

					if math.Abs(float64(report[i]-report[i+1])) > 3 {
							isSubtle = false
					}
			}

			return (isDecreasing || isIncreasing) && isSubtle
	}

	if checkSafe(report) {
			return true
	}

	for i := range report {
			modifiedReport := append([]int{}, report[:i]...)
			modifiedReport = append(modifiedReport, report[i+1:]...)
			if checkSafe(modifiedReport) {
					return true
			}
	}

	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	reports := make([][]int, 0)

	for scanner.Scan() {
		reports = appendReports(reports, scanner.Text())
	}

	counter := 0
	for _, report := range reports {
		if isSafe(report) {
			counter++
		}
	}

	fmt.Println(counter)
}