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

		//do the math 2*l*w + 2*w*h + 2*h*l + the smallest side again
		dims := []int{}
		dims = append(dims, ints[0]*ints[1])
		dims = append(dims, ints[0]*ints[2])
		dims = append(dims, ints[1]*ints[2])

		sort.Ints(dims)

		//add it all up
		packageSum := 2*dims[0] + 2*dims[1] + 2*dims[2] + dims[0]

		total += packageSum

	}
	fmt.Println(total)
}
