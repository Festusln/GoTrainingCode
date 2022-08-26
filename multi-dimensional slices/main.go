package main

import "fmt"

func main() {
	// amounts := [][]int{
	// 	{100, 200, 300},
	// 	{250},
	// 	{500, 400},
	// }
	amount := make([][]int, 0, 5)
	amount = append(amount, []int{30, 80, 50})
	amount = append(amount, []int{400, 300, 60})

	for i, amount := range amount {
		var total int
		for _, spending := range amount {
			total += spending
		}

		fmt.Printf("Day %d: %d\n", i, total)
	}
}
