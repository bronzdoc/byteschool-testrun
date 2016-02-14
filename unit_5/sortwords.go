package main

import "os"
import "io/ioutil"
import "fmt"

const (
	inputFilename  = "words.txt"
	outputFilename = "sorted_words.txt"
)

func main() {
	data, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(data)

	words := splitLines(str)
	// assume words.txt has one word per line
	sortStrings(words)
	str = joinLines(words)

	err = ioutil.WriteFile(outputFilename, []byte(str), os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// splits a string into substrings using newlines as separators
// e.g. for "foo\nbar\nack" returns []string{"foo", "bar", "ack"}
func splitLines(s string) []string {
	var str_a []string
	str := ""
	for _, char := range s {
		if char != '\n' {
			str += string(char)
		} else {
			str_a = append(str_a, string(str))
			str = ""
		}
	}

	fmt.Printf(str)
	return str_a
}

// Sorts the strings of the slice alphabetically.
// We compare strings with the > operator. A > comparison of strings is not
// exactly alphabetical, but it will serve our purposes.
// (uppercase letters have lower number values than lowercase letter
// in ASCII/Unicode and so will sort in front,
// e.g. "Zebra" sorts before "apple")
func sortStrings(strs []string) {
	// bubble sort
	for i := 0; i < len(strs)-1; i++ {
		for j := 0; j < len(strs)-1-i; j++ {
			a, b := strs[j], strs[j+1]
			if a > b {
				strs[j+1], strs[j] = a, b // swap
			}
		}
	}
}

func joinLines(strs []string) string {
	str := ""
	for _, s := range strs {
		str += (s+"\n")
	}
	return str
}
