package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("A value is expected to be entered")
		return
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid entries:", err)
		return
	}
	fmt.Println("The entered value convereted to int is:", number)
}
