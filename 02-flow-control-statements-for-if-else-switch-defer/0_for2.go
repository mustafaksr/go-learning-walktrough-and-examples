package main

import "fmt"

func main() {
	sum := 1
	for sum < 10000 {
		sum += sum
		fmt.Printf("i value is %v\n", sum)

	}
	fmt.Printf("sum of all i is %v", sum)

}
