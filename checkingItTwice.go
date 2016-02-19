package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := os.Args[1]

	file, _ := os.Open(filePath)

	defer file.Close()

	reader := bufio.NewReader(file)

	scanner := bufio.NewScanner(reader)

	scanner.Split(bufio.ScanLines)

	nice := 0

	for scanner.Scan() {
		currentString := scanner.Text()

		//verify xyx character pattern
		if hasXYXSequence(currentString) {
			//verify pair of pairs
			if hasPairOfPairs(currentString) {
				nice++
			}
		}
	}
	fmt.Println(nice)
}

func hasPairOfPairs(str string) bool {
	runeSlice := []rune(str)
	n := len(runeSlice)
	for i := 0; i < n-2; i++ {
		c1 := runeSlice[i]
		for j := i + 2; j < n-1; j++ {
			if c1 == runeSlice[j] && runeSlice[i+1] == runeSlice[j+1] {
				fmt.Println(str)
				return true
			}
		}
	}
	return false
}

func hasXYXSequence(str string) bool {
	runeSlice := []rune(str)
	n := len(runeSlice)
	for i := 0; i < n-2; i++ {
		if runeSlice[i] == runeSlice[i+2] {

			return true
		}
	}
	return false
}
