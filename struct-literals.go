package main

import "fmt"

type Vertex struct {
	X int
	Y string
}

var (
	v1 = Vertex{1, "Hello"}
	v2 = Vertex{X: 1} // x = 1 y = ""
	v3 = Vertex{}     // x=0 y = ""
	p  = &Vertex{1, "world"}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
