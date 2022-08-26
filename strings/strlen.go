package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	arg := os.Args[1]
	len := len(arg)
	fmt.Println("The length of the string using len is ", len)
	len = utf8.RuneCountInString(arg)
	fmt.Println("The length of the string using rune is ", len)
	fmt.Println(strings.Repeat("!", len) + strings.ToUpper(arg) + strings.Repeat("!", len)) //to avoid using the repeat function twice
	s := strings.Repeat("!", len)
	fmt.Println(s + strings.ToUpper(arg) + s)
	fmt.Println("The upper case of the entered string is ", strings.ToLower(arg))
}
