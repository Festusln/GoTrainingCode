package main

import "fmt"

type milli float64

type meters float64

func main() {
	var mm milli
	var m meters = 1000
	mm = milli(m) * 0.3124
	fmt.Println(mm)
}
