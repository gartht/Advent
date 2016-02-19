package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
		match1, _ := regexp.MatchString(`ab|cd|pq|xy`, currentString)
		if match1 {
			continue
		}

		if !hasRuneSequence(currentString) {
			continue
		}

		match2, _ := regexp.MatchString(`([aeiou].*){3,}`, currentString)
		if !match2 {
			continue
		}
		fmt.Println(currentString)
		nice++
	}

	fmt.Println(nice)
}

func hasRuneSequence(str string) bool {
	runeSlice := []rune(str)
	n := len(runeSlice)
	for i := 0; i < n-1; i++ {
		if runeSlice[i] == runeSlice[i+1] {
			return true
		}
	}
	return false
}
