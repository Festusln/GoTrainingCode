package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	ft, _ := strconv.ParseFloat(arg, 64)
	meters := ft * 0.3048
	fmt.Printf("The equivalent value of %sft in meter is %.2f\n", arg, meters)
}
