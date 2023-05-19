package main

import (
	"fmt"
	"math"
)

var var1 bool

// bool
// string
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // alias for uint8
// rune // alias for int32 unicode point
// float32 float64
// complex64 complex128

func main() {
	var i int
	var1 = true
	i = 10
	j := "ashish"
	fmt.Println(var1, i)
	fmt.Printf("Type: %T, Value:%v\n", j, j)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
