package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.0234234, -74.123441,
	}
	fmt.Println(m["Bell Labs"])

	m1 := make(map[string]int)
	m1["Answer"] = 42
	fmt.Println("The value is: ", m1["Answer"])

	m1["Answer"] = 48
	fmt.Println("The value is: ", m1["Answer"])

	delete(m1, "Answer")
	fmt.Println("The value is: ", m1["Answer"])

	v, ok := m1["Answer"]
	fmt.Println("The value: ", v, "Present?", ok)
	if ok {
		fmt.Println("the value:", v)
	} else {
		fmt.Println("invalid key")
	}

}
