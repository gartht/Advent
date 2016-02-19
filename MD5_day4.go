package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
)

func main() {

	i := 254575

	for {
		testString := "bgvyzdsv" + strconv.Itoa(i)
		h := md5.New()

		io.WriteString(h, testString)
		checkStr := hex.EncodeToString(h.Sum(nil))
		//hash := md5.Sum(testString)

		if checkStr[0:6] == "000000" {
			fmt.Println(i)
			break
		}
		i++
	}

	validStr := "abcdef609043"
	h := md5.New()
	io.WriteString(h, validStr)
	fmt.Println(hex.EncodeToString(h.Sum(nil))[0:5])
}
