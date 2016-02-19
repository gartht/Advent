package main

import (
	"bufio"
	"fmt"
	"os"
)

var hIndex int
var vIndex int

func main() {
	filePath := os.Args[1]

	file, _ := os.Open(filePath)

	defer file.Close()

	reader := bufio.NewReader(file)

	scanner := bufio.NewScanner(reader)

	deliveryMap := make(map[int]map[int]int)

	//deliveryMap := make([]map[int]int{})

	scanner.Split(bufio.ScanRunes)

	deliveryMap[0] = make(map[int]int)
	deliveryMap[0][0] = 1

	hIndex = 0
	vIndex = 0

	for scanner.Scan() {
		//parse out the data using the x as the seperator
		changeIndices(scanner.Text())
		if deliveryMap[hIndex] == nil {
			deliveryMap[hIndex] = make(map[int]int)
		}

		deliveryMap[hIndex][vIndex]++

	}

	houseCount := 0

	for _, v := range deliveryMap {
		houseCount += len(v)
	}
	fmt.Println(houseCount)
}

func changeIndices(token string) {
	switch token {
	case ">":
		hIndex++

	case "<":
		hIndex--

	case "^":
		vIndex++

	case "v":
		vIndex--
	}
}
