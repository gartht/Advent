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

	runeMap := make(map[string]int)

	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		runeMap[scanner.Text()] += 1
	}

	up := runeMap["("]
	down := runeMap[")"]

	fmt.Println(up - down)

}
