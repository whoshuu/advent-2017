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
	halfway := (len(input) - 1) / 2
	for i := 0; i < halfway; i++ {
		total += checkDuplicate(input[i], input[i+halfway])
	}

	fmt.Println(total * 2)
}
