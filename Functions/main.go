package main

import (
	"fmt"
	"math"
)

func calculate(x, y int) (add, sub, multiply int) {
	add = x + y
	sub = x - y
	multiply = x * y
	return
}

func main() {
	plus, minus, times := calculate(5, 6)
	fmt.Println("Multiplication gives", times)
	fmt.Println("Addition gives", plus)
	fmt.Println("Subtraction give", minus)
	math.Sqrt(9)
}
