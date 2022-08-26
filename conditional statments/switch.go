package main

import (
	"fmt"
)

func main() {
	i := 30
	switch {
	case i > 20:
		fmt.Print("Value is greater than 20 ")
		// fallthrough
	case i >= 0:
		fmt.Print("Value is less than 20 ")
		// fallthrough
	default:
		fmt.Print("Value is negative")
	}
}
