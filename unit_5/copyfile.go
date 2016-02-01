package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	inputFilename  = "words.txt"
	outputFilename = "results.txt"
)

func main() {
	data, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(data)
	ioutil.WriteFile(outputFilename, []byte(str), os.ModePerm)
}
