package main

import (
	"fmt"
	"os"
	"strconv"
)

func divide(x, y int, x2, y2 float64) float64 {
	return float64(x+y) / (x2 + y2)
}

func main() {
	if len(os.Args) != 5 {
		fmt.Println("Usage: go run main.go <int> <int> <float> <float>")
		return
	}

	arg1, err1 := strconv.Atoi(os.Args[1])
	arg2, err2 := strconv.Atoi(os.Args[2])
	arg3, err3 := strconv.ParseFloat(os.Args[3], 64)
	arg4, err4 := strconv.ParseFloat(os.Args[4], 64)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		fmt.Println("Invalid arguments")
		return
	}

	result := divide(arg1, arg2, arg3, arg4)
	fmt.Println("Result:", result)
}
