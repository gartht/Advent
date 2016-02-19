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

	scanner.Split(bufio.ScanRunes)

	currentFloor := 0

	pos := 0

	for scanner.Scan() {
		char := scanner.Text()
		if char == "(" {
			currentFloor += 1
		} else if char == ")" {
			currentFloor -= 1
		}
		pos++
		if currentFloor == -1 {
			fmt.Println(pos)
			break
		}
	}

}
