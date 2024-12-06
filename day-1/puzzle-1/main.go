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

func appendIdToList(list []int ,txt string) []int {
	value, err := strconv.Atoi(txt)
	if (err != nil) {
		panic("Not a number")
	}

	return append(list, value)
}

func main(){
	file, err := os.Open("input.txt")

	if (err != nil) {
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

	for i := range(leftList) {
		distance := leftList[i] - rightList[i]
		result += max(distance, -distance)
	}

	fmt.Println("Result:", result)
}