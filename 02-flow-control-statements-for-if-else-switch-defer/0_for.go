package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Printf("i value is %v\n", i)

	}
	fmt.Printf("sum of all i is %v", sum)

}
