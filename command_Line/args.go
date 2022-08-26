package main

import (
	"fmt"
	"os"
)

// func main() {
// 	fmt.Printf("The path name: %#v\n", os.Args[0]) //this print the value and type. The type is string. If println is used, it only prints the value which will be printed as array
// 	fmt.Printf("First argument: %#v\n", os.Args[1])
// 	fmt.Printf("Second aergument: %#v\n", os.Args[2])
// 	fmt.Printf("The length of argument: %d\n", len(os.Args))
// }
func main() {
	name := os.Args[1]
	fmt.Println("The path name is ", os.Args[0])
	fmt.Println("Welcome ", name)
	fmt.Println("Hope you had a wonderful day")
}
