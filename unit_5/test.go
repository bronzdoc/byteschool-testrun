package main

import "fmt"

func PrintPrimes() {
	first_ten_primes := [10]int{2, 3, 5, 7, 9, 11, 13, 17, 19, 23}
	for i := 0; i < len(first_ten_primes); i += 1 {
		fmt.Println(first_ten_primes[i])
	}
}

func main() {
	PrintPrimes()
}
