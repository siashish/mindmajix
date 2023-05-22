package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	fmt.Println(len(q))
	fmt.Println(cap(q))

	q = append(q, 17)

	arr := [10]int{10, 20}
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
}
