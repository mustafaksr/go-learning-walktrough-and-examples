package main

import "fmt"

var pow [3][3]int

func main() {
	ii := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("2**%d = %d\n", i, i*i)
			pow[i][j] = ii * ii
			ii += 1
		}
	}
	fmt.Println(pow)
}
