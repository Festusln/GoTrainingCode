package main

import "fmt"

func main() {
	items := []string{"akinkunmi", "festus", "alamu", "racheal", "oluwakemi",
		"akanmu", "olatunde", "jeremiah", "adewole", "JESUS"}
	sliceSize := len(items)
	fmt.Println("Size of the slice is ", sliceSize)
	const numberPerPage = 3

	for currentNumber := 0; currentNumber < sliceSize; {
		if (currentNumber + numberPerPage) > sliceSize {
			fmt.Println(items[currentNumber:])
			// return
		} else {
			fmt.Println(items[currentNumber : currentNumber+numberPerPage])
		}

		currentNumber += (currentNumber + numberPerPage)
	}
}
