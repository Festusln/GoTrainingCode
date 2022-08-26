package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var num float64 = 52
	var result = math.Sqrt(num)
	fmt.Println("The square root of", num, "is:", result)
	calc()
	fmt.Println("Today is ", time.Now().Weekday())
}
