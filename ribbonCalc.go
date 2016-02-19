package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	filePath := os.Args[1]

	file, _ := os.Open(filePath)

	defer file.Close()

	reader := bufio.NewReader(file)

	scanner := bufio.NewScanner(reader)

	scanner.Split(bufio.ScanLines)

	total := 0

	for scanner.Scan() {
		//parse out the data using the x as the seperator
		lhw := strings.Split(scanner.Text(), "x")

		ints := []int{}

		//convert it to numbers
		for _, v := range lhw {
			x, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			ints = append(ints, x)
		}

		sort.Ints(ints)

		baseRibbon := 2*ints[0] + 2*ints[1]
		bow := ints[0] * ints[1] * ints[2]

		total += baseRibbon + bow

	}
	fmt.Println(total)
}
