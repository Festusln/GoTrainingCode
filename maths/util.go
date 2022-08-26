package main

import (
	"fmt"
	"math"
)

func calc() {
	var num = 4.578692
	fmt.Printf("The ceil of %f is %.3f\n", num, math.Ceil(num))
	fmt.Printf("The floor of %f is %.3f\n", num, math.Floor(num))
	fmt.Printf("Rounding off of %f is %.3f\n", num, math.Round(num))
}
