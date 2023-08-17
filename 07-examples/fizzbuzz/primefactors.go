package main

import (
	"fmt"
	"os"
	"strconv"
)

func primeFactors(n int) []int {
	factors := []int{}

	// Divide by 2 until it's an odd number
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}

	// Divide by odd numbers starting from 3
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	// If n is a prime greater than 2
	if n > 2 {
		factors = append(factors, n)
	}

	return factors
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <number>")
		return
	}

	numberStr := os.Args[1]
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		fmt.Println("Invalid number:", numberStr)
		return
	}

	factors := primeFactors(number)
	fmt.Printf("Prime factors of %d are: %v\n", number, factors)
}
