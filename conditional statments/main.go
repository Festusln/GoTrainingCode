package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No value supplied")
		return
	}
	argc := os.Args[1]
	km, _ := strconv.ParseFloat(argc, 64)
	m := km * 1000
	fmt.Println("The equivalent value entered in meter is ", m)
}
