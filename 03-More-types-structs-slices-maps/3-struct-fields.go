package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type ZCoordinates struct {
	Z0 int
	Z1 int
}

type ThreeDimPoint struct {
	X int
	Y int
	Z ZCoordinates
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	t := ThreeDimPoint{
		3,
		5,
		ZCoordinates{
			Z0: 30,
			Z1: 40,
		}}

	fmt.Println(v.X)
	fmt.Println(t.Z)
	fmt.Println(t.Z.Z0)
}
