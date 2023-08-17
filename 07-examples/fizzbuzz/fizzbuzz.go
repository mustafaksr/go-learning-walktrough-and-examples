package main

import (
	"fmt"
	"os"
	"strconv"
)

func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
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

	fizzBuzz(number)
}
