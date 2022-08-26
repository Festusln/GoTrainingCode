package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	list = "festus racheal olatunde olumide nathaniel akinkunmi oladele jonathan, festus"
)

func main() {
	var count int
	word := os.Args
	if len(word) < 2 {
		fmt.Println("Supply name(s) to search")
		return
	}
	lists := strings.Fields(list)

	for _, w := range word {

		for _, l := range lists {
			if strings.ToLower(l) == strings.ToLower(w) {
				fmt.Printf("Found %q\n", w)
				count++
				break
			}
		}

	}
	if count < 1 {
		fmt.Println("No search word found")
	}

}
