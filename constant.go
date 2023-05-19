package main

import "fmt"

const pi = 3.14

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	const name = "Ashish"
	fmt.Println("hello", name)
	fmt.Println("happy", pi)

	const truth = true
	fmt.Println("Go rule ", truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
