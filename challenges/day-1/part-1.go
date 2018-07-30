package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func checkDuplicate(a, b byte) int {
	if a == b {
		duplicate, err := strconv.Atoi(string(a))
		if err != nil {
			panic(err)
		}
		return duplicate
	}
	return 0
}

func main() {
	inputFile, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	input, err := ioutil.ReadAll(inputFile)
	if err != nil {
		panic(err)
	}

	total := 0

	for i := 0; i < len(input)-2; i++ {
		total += checkDuplicate(input[i], input[i+1])
	}
	if input[0] == input[len(input)-2] {
		total += checkDuplicate(input[0], input[len(input)-2])
	}

	fmt.Println(total)
}
