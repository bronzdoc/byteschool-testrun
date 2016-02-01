package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)

func main() {
	numUint, err := parseUint("952234251")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The parsed uint is", numUint)

	numInt, err := parseInt("-241")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The parsed int is", numInt)

	numInt, err = parseInt("241")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The parsed int is", numInt)
}

// return non-nil error when string includes non-numeric characters
func parseUint(str string) (uint, error) {
	var digitsByRune = map[rune]uint{
		48: 0,
		49: 1,
		50: 2,
		51: 3,
		52: 4,
		53: 5,
		54: 6,
		55: 7,
		56: 8,
		57: 9,
	}

	// create []rune of the string's characters so that we can iterate through the runes backwards
	digits := make([]uint, utf8.RuneCountInString(str)) // returns count of characters in the string
	for i, ch := range str {
		digit, exists := digitsByRune[ch]
		if !exists {
			return 0, fmt.Errorf("String does not represent a valid uint.")
		}
		digits[i] = digit
	}

	result := uint(0)
	n := len(digits) - 1
	for _, num := range digits {
		result += uint(math.Pow(10, float64(n))) * num
		n -= 1
	}
	return result, nil
}

// return non-nil error when string includes non-numeric characters (other than a leading "-")
func parseInt(str string) (int, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("String does not represent a valid uint.")
	}
	return num, err
}
