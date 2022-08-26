package main

import (
	"fmt"
	"time"
)

func main() {
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("Good Morning")
	case hour >= 12 && hour < 16:
		fmt.Println("Good Afternoon")
	case hour >= 16 && hour < 21:
		fmt.Println("Good evening")
	default:
		fmt.Println("Goodnight")
	}
}
