package main

import (
	"fmt"
	"os"
)

func swap(x, y string) (string, string) { return y, x }

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <string1> <string2>")
		return
	}

	a, b := swap(os.Args[1], os.Args[2])
	fmt.Println(a, b)
}
