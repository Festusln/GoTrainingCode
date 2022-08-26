package main

import "fmt"

func check() {
	nums := 2
	if nums > 1 {
		fmt.Println("Value is more than one")
	} else if nums == 1 {
		fmt.Println("Value is one")
	} else {
		fmt.Println("Value is less than one")
	}
}
