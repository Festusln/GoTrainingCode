//To get randon no, we use the "package math/rand"
//There are different functions for generating random no in the package
//The one to use depend on the type of random no we want to generate
//E.g: to generate float64, we use func Float64() float64
//func Int() int :this will produce any random numbers
//func Intn(n int) int :this will produce random number within range 0 to n without including n but t includes 0

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
// 	guess := 10
// 	for i := 0; i < guess; i++ {
// 		fmt.Printf("%d ", rand.Intn(10)) //generating rand no from 0 to 9
// 	}
// 	fmt.Println()

// } //no matter how much time this is run, it will continue to give the same random numbers
//Tosolve thsi problem, we need to seed the random number. each seed will prodce diff rand no.
//So far time changes constantly, we use time to seed the random no

func main() {
	rand.Seed(time.Now().UnixNano())
	guess := 10
	for i := 0; i < guess; i++ {
		fmt.Printf("%d ", rand.Intn(10))
	}
	fmt.Println()
}
