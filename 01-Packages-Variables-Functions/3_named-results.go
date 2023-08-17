package main

import (
	"fmt"
)

func split(sum int) (x, y float32) {

	x = float32(sum) * 4 / 9
	y = float32(sum) - x

	return //naked return
}

func main() {
	fmt.Println(split(17))
}
