package main

import "fmt"

type Vertex struct {
	X int
	Y string
}

func main() {
	fmt.Println(Vertex{1, "Ashish"})

	v := Vertex{1, "Ashish"}
	v.X = 4
	v.Y = "Singh"
	fmt.Println(v.X)
	fmt.Println(v.Y)

	p := &v
	p.X = 1e9
	fmt.Println(v)

}
