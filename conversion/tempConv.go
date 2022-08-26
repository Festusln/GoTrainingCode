//Converting fahrenheit to celsius by accepting value through command line
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	fahrenheit, _ := strconv.ParseFloat(arg, 64)
	celsius := (fahrenheit - 32) * (float64(5) / float64(9))
	fmt.Printf("The value of %.2fF in celsius is %.2f\n", fahrenheit, celsius)
}
