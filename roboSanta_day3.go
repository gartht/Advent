package main

import (
	"bufio"
	"fmt"
	"os"
)

var santaHindex int
var santaVindex int
var roboHindex int
var roboVindex int

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

	santaHindex = 0
	santaVindex = 0
	roboHindex = 0
	roboVindex = 0

	scanPosition := 0

	for scanner.Scan() {
		//parse out the data using the x as the seperator
		if scanPosition%2 == 0 {
			changeIndices(scanner.Text(), &santaHindex, &santaVindex)
			updateDeliveryMap(deliveryMap, santaHindex, santaVindex)
		} else {
			changeIndices(scanner.Text(), &roboHindex, &roboVindex)
			updateDeliveryMap(deliveryMap, roboHindex, roboVindex)
		}
		scanPosition++
	}

	houseCount := 0

	for _, v := range deliveryMap {
		houseCount += len(v)
	}
	fmt.Println(houseCount)
}

func updateDeliveryMap(deliveryMap map[int]map[int]int, hIndex int, vIndex int) {
	if deliveryMap[hIndex] == nil {
		deliveryMap[hIndex] = make(map[int]int)
	}

	deliveryMap[hIndex][vIndex]++
}

func changeIndices(token string, hIndex *int, vIndex *int) {
	switch token {
	case ">":
		*hIndex++

	case "<":
		*hIndex--

	case "^":
		*vIndex++

	case "v":
		*vIndex--
	}
}
