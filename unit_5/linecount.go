package main

import "io/ioutil"
import "fmt"

const inputFilename = "words.txt"

func main() {
	data, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(data)
	wordCount := 0
	for _, char := range str {
		if char == '\n' {
			wordCount += 1
		}
	}
	fmt.Println(wordCount)
}
