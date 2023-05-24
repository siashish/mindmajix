package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type myFloat float64

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f myFloat) Absfloat() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	f := myFloat(-math.Sqrt2)
	fmt.Println(f.Absfloat())
}
