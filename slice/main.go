package main

import (
	"fmt"
	"unsafe"
)

func main() {
	message := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Printf("%s\n", message[:])
	var name int
	fmt.Println(unsafe.Sizeof(name))

}
