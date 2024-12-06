package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func appendIdToList(list []int, id string) []int {
	result, err := strconv.Atoi(id)
	if (err != nil) {
		panic("Not a number")
	}

	return append(list, result)
}


func main() {
	file, err := os.Open("input.txt")

	if(err != nil) {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var leftList []int
	var rightList []int
	
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		
		leftList = appendIdToList(leftList, line[0])
		rightList = appendIdToList(rightList, line[1])
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	result := 0
	
	for _, l := range(leftList) {
			count := 0
			for _, r := range(rightList) {
				if (l == r) {
					count++
				}
			}
			result += l * count
	}

	fmt.Println("Result: ", result)
}