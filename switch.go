package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X(Mac)")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s \n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Friday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In Two days")
	default:
		fmt.Println("Too far away")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}
}
