//printing words in lines with max allowable char
//without cutting any of the words

package main

import (
	"fmt"
	"unicode"
)

func main() {
	const words = "Nigeria is a country located in western Africa, it is sorrounded by many countries, boardered with Cameroon in the east, in the north by Niger Republic, in the south by Atlantic ocean"

	const maxWidth = 10
	len := 0
	for _, c := range words {
		fmt.Printf("%c", c)
		if len++; len >= maxWidth && c != '\n' && unicode.IsSpace(c) {
			fmt.Println()
			len = 0
		}
	}
}
